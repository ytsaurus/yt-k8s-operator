package ytconfig

import (
	"k8s.io/apimachinery/pkg/types"

	ytv1 "github.com/ytsaurus/ytsaurus-k8s-operator/api/v1"
)

type NodeGenerator struct {
	BaseGenerator
}

func NewRemoteNodeGenerator(
	key types.NamespacedName,
	clusterDomain string,
	commonSpec ytv1.CommonSpec,
	masterConnectionSpec ytv1.MasterConnectionSpec,
	masterCachesSpec *ytv1.MasterCachesSpec,
) *NodeGenerator {
	baseGenerator := NewRemoteBaseGenerator(
		key,
		clusterDomain,
		commonSpec,
		masterConnectionSpec,
		masterCachesSpec,
	)
	return &NodeGenerator{
		BaseGenerator: *baseGenerator,
	}
}

func NewLocalNodeGenerator(
	ytsaurus *ytv1.Ytsaurus,
	clusterDomain string,
) *NodeGenerator {
	baseGenerator := NewLocalBaseGenerator(ytsaurus, clusterDomain)
	return &NodeGenerator{
		BaseGenerator: *baseGenerator,
	}
}
