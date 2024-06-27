package ytconfig

type UIAuthenticationType string

const (
	uiAuthenticationBasic UIAuthenticationType = "basic"
)

type UIPrimaryMaster struct {
	CellTag int16 `yson:"cellTag"`
}

type UICluster struct {
	ID             string               `yson:"id"`
	Name           string               `yson:"name"`
	Proxy          string               `yson:"proxy"`
	Secure         bool                 `yson:"secure"`
	Authentication UIAuthenticationType `yson:"authentication"`
	Group          string               `yson:"group"`
	Theme          string               `yson:"theme"`
	Environment    string               `yson:"environment"`
	Description    string               `yson:"description"`
	PrimaryMaster  UIPrimaryMaster      `yson:"primaryMaster"`
}

type UIClusters struct {
	Clusters []UICluster `yson:"clusters"`
}

func getUIClusterCarcass() UICluster {
	return UICluster{
		Secure:         false,
		Authentication: uiAuthenticationBasic,
		Group:          "My YTsaurus clusters",
		Description:    "My first YTsaurus. Handle with care.",
	}
}

type UICustomSettings struct {
	DirectDownload *bool `yson:"directDownload,omitempty"`
}

type UICustom struct {
	OdinBaseUrl *string           `yson:"odinBaseUrl,omitempty"`
	Settings    *UICustomSettings `yson:"uiSettings,omitempty"`
}
