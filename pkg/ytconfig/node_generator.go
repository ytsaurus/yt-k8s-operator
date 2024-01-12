package ytconfig

import (
	"k8s.io/apimachinery/pkg/types"

	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
)

type NodeGenerator struct {
	BaseGenerator
}

func NewRemoteNodeGenerator(
	key types.NamespacedName,
	clusterDomain string,
	configSpec ytv1.ConfigurationSpec,
	masterConnectionSpec ytv1.MasterConnectionSpec,
) *NodeGenerator {
	baseGenerator := NewRemoteBaseGenerator(
		key,
		clusterDomain,
		configSpec,
		masterConnectionSpec,
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
