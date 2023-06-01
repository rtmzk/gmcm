package v1

import (
	"gmcm/pkg/utils/devices"
	"gmcm/pkg/utils/storage"
)

type ConnectionCheckParams struct {
	IP          []string `json:"ip"`
	SSHPort     string   `json:"ssh_port,omitempty"`
	SSHUser     string   `json:"ssh_user"`
	SSHPassword string   `json:"ssh_password,omitempty"`
	NoPass      bool     `json:"is_no_pass,omitempty"`
	NodeRole    string   `json:"-"`
}

type Hosts struct {
	//NodeType       string `json:"node_type" gorm:"column:type;default:storage"`
	ObjectMeta  `json:"metadata,omitempty"`
	IP          string `json:"ip" gorm:"column:ip;not null"`
	NodeRole    string `json:"node_role" gorm:"column:role;not null"`
	SSHUser     string `json:"ssh_user" gorm:"column:sshUser;not null"`
	SSHPort     string `json:"ssh_port" gorm:"column:sshPort"`
	SSHPassword string `json:"ssh_password" gorm:"column:sshPass"`
	NoPass      bool   `json:"is_no_pass,omitempty" gorm:"column:nopass"`
}

// Status 标识存储初始化状态
// 0: 未初始化
// 1: 已初始化
type Status struct {
	ObjectMeta `json:"-"`
	InitStatus int `json:"status" gorm:"column:status;default:0"`
}

type StorageNetworks struct {
	PublicNetwork  string `json:"public_network"`
	ClusterNetwork string `json:"cluster_network,omitempty"`
}

type Devices struct {
	IP     string            `json:"ip"`
	Device []devices.Devices `json:"device"`
}

type DeviceList struct {
	StorageNetworks
	Replicas int       `json:"replicas"`
	Devices  []Devices `json:"devices"`
}

type HostsList struct {
	Hosts []Hosts `json:"hosts"`
}

type FlushTemplate struct {
	Pools          storage.Pools
	PublicNetwork  string
	ClusterNetwork string
	Device         map[string][]devices.Devices
	Host           []Hosts
	NoSSD          bool
}

type InstallLogs struct {
	Offset int64  `json:"offset"`
	Text   string `json:"text"`
}

type Rule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Func        string `json:"func"`
	ScriptOut
}

type CheckRules struct {
	Rules []Rule `json:"rules"`
}

type ScriptOut struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
