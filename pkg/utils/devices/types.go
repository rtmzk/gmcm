package devices

const (
	HDD                     DeviceType = "hdd"
	SSD                     DeviceType = "ssd"
	DeviceDefaultPathPrefix string     = "/dev/"
)

type Devices struct {
	Name        string     `json:"name"`
	Size        string     `json:"size"`
	Type        DeviceType `json:"type"`
	DefaultType DeviceType `json:"default_type"`
	Enabled     bool       `json:"enabled"`
	Cached      bool       `json:"cached"`
}

type DeviceType string

var DeviceTypeMapping = map[DeviceType]DeviceType{
	"1": HDD,
	"0": SSD,
}
