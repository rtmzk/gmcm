package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
	"gmcm/pkg/utils/devices"
	"strings"
)

func GetHostDevices(c *gin.Context) {
	var allHosts []v1.Hosts
	var dev v1.Devices
	var devs []v1.Devices
	var replicas int
	if err := db.Client().Table("hosts").Find(&allHosts).Error; err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "Could not get all host."), nil)
		return
	}
	for _, h := range allHosts {
		if !strings.Contains(h.NodeRole, "osd") {
			continue
		}
		cli, err := newSSHClient(h, h.IP)
		if err != nil {
			continue
		}

		d, err := devices.GetRemoteDevices(cli)
		if err != nil {
			core.WriteResponse(c, errors.WithCode(code.ErrUnknown, "Unknown"), nil)
			return
		}

		if len(d) == 0 {
			continue
		}

		dev.IP = h.IP
		dev.Device = d

		devs = append(devs, dev)
	}

	replicas = len(devs)
	if len(devs) > 3 {
		replicas = 3
	}

	devsList := &v1.DeviceList{
		Replicas: replicas,
		Devices:  devs,
	}

	core.WriteResponse(c, nil, &devsList)
}
