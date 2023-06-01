package storage

type Pools struct {
	Sp storePool
}
type storePool struct {
	Root     int `json:"_rgw_root"`
	Data     int `json:"master_rgw_buckets_data"`
	Index    int `json:"master_rgw_buckets_index"`
	Control  int `json:"master_rgw_control"`
	Meta     int `json:"master_rgw_meta"`
	Log      int `json:"master_rgw_log"`
	NonEc    int `json:"master_rgw_buckets_non-ec"`
	Replicas int `json:"replicas"`
}

func GetDefalutPoolSize() Pools {
	var defaultPoolSize Pools

	defaultPoolSize.Sp.Root = 8
	defaultPoolSize.Sp.Data = 64
	defaultPoolSize.Sp.Index = 32
	defaultPoolSize.Sp.Control = 16
	defaultPoolSize.Sp.Meta = 16
	defaultPoolSize.Sp.Log = 16
	defaultPoolSize.Sp.NonEc = 32

	return defaultPoolSize
}
