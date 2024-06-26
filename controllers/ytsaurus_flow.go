package controllers

import (
	"context"
	"errors"
	"fmt"

	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
	"github.com/ytsaurus/yt-k8s-operator/pkg/components"
	"github.com/ytsaurus/yt-k8s-operator/pkg/consts"
)

func getStatuses(
	ctx context.Context,
	registry *componentRegistry,
	order [][]consts.ComponentType,
) (map[string]components.ComponentStatus, error) {
	statuses := make(map[string]components.ComponentStatus)
	for _, batch := range order {
		batchComps := registry.listByType(batch...)
		for _, c := range batchComps {
			componentStatus, err := c.Status(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to get component %s status: %w", c.GetName(), err)
			}
			statuses[c.GetName()] = componentStatus
		}
	}
	return statuses, nil
}

// componentsOrder is an order in which components will be built.
// batches of components are supported, but since we have status update conflict
// almost every second reconciliation, we start with truly linear flow: one component at a time.
var componentsOrder = [][]consts.ComponentType{
	// This is not declared here, but
	// at first, we check if master is *built* (not updated) before everything else.
	// batch #1
	{consts.YtsaurusClientType},
	{consts.DiscoveryType},
	{consts.HttpProxyType},
	{consts.RpcProxyType},
	{consts.TcpProxyType},
	{consts.DataNodeType},
	{consts.ExecNodeType},
	{consts.MasterCacheType},

	// batch #2
	{consts.TabletNodeType},
	{consts.UIType},
	{consts.ControllerAgentType},
	{consts.YqlAgentType},

	// batch #3
	{consts.SchedulerType},
	{consts.QueryTrackerType},
	{consts.QueueAgentType},

	// batch #4
	{consts.StrawberryControllerType},

	// batch #5
	{
		// Here we UPDATE master after all the components, because it shouldn't be newer
		// than others.
		// Currently, we guarantee that only for the case when components are not redefine their images.
		consts.MasterType,
	},
}

func syncComponents(
	ctx context.Context,
	registry *componentRegistry,
	resource *ytv1.Ytsaurus,
) (components.ComponentStatus, error) {
	statuses, err := getStatuses(ctx, registry, componentsOrder)
	if err != nil {
		return components.ComponentStatus{}, err
	}
	if err = logComponentStatuses(ctx, registry, statuses, componentsOrder, resource); err != nil {
		return components.ComponentStatus{}, err
	}

	// Special check before everything other component (including master) update.
	masterBuildStatus, err := getStatusForMasterBuild(ctx, registry.master)
	if err != nil {
		return components.ComponentStatus{}, err
	}
	switch masterBuildStatus.SyncStatus {
	case components.SyncStatusBlocked:
		return masterBuildStatus, nil
	case components.SyncStatusNeedSync:
		return masterBuildStatus, registry.master.BuildInitial(ctx)
	}

	var batchToSync []component
	for _, typesInBatch := range componentsOrder {
		compsInBatch := registry.listByType(typesInBatch...)
		for _, comp := range compsInBatch {
			status := statuses[comp.GetName()]
			if status.SyncStatus != components.SyncStatusReady && batchToSync == nil {
				batchToSync = compsInBatch
			}
		}
	}

	if batchToSync == nil {
		// YTsaurus is running and happy.
		return components.ComponentStatus{SyncStatus: components.SyncStatusReady}, nil
	}

	// Run sync for non-ready components in the batch.
	batchNotReadyStatuses := make(map[string]components.ComponentStatus)
	var errList []error
	for _, comp := range batchToSync {
		status := statuses[comp.GetName()]
		if status.SyncStatus == components.SyncStatusReady {
			continue
		}
		batchNotReadyStatuses[comp.GetName()] = status
		if err = comp.Sync(ctx); err != nil {
			errList = append(errList, fmt.Errorf("failed to sync %s: %w", comp.GetName(), err))
		}
	}

	if len(errList) != 0 {
		return components.ComponentStatus{}, errors.Join(errList...)
	}

	// Choosing the most important status for the batch to report up.
	batchStatus := components.ComponentStatus{
		SyncStatus: components.SyncStatusUpdating,
		Message:    "",
		Stage:      "",
	}
	for compName, st := range batchNotReadyStatuses {
		if st.SyncStatus == components.SyncStatusBlocked {
			batchStatus.SyncStatus = components.SyncStatusBlocked
		}
		if compName == registry.master.GetName() {
			// TODO: add scheduler (and maybe query tracker if they will be in separate batches)
			batchStatus.Stage = st.Stage
		}
		batchStatus.Message += fmt.Sprintf("; %s=%s (%s)", compName, st.SyncStatus, st.Message)
	}
	return batchStatus, nil
}

func getStatusForMasterBuild(ctx context.Context, master masterComponent) (components.ComponentStatus, error) {
	masterBuiltInitially, err := master.IsBuildInitially(ctx)
	if err != nil {
		return components.ComponentStatus{}, err
	}
	masterNeedBuild, err := master.NeedBuild(ctx)
	if err != nil {
		return components.ComponentStatus{}, err
	}
	masterRebuildStarted := master.IsRebuildStarted()

	if !masterBuiltInitially {
		// This only happens once on cluster initialization.
		return components.NeedSyncStatus("master initial build"), nil
	}

	if masterNeedBuild && !masterRebuildStarted {
		// Not all the master's sub-resources are running, and it is NOT because master is in update stage
		// (in which is reasonable to expect some not-yet-built sub-resources).
		// So we can't proceed with update, because almost every component need working master to be updated properly.
		return components.ComponentStatus{
			SyncStatus: components.SyncStatusBlocked,
			Message:    "Master is not built, cluster can't start the update",
		}, nil
	}
	return components.ReadyStatus(), nil
}
