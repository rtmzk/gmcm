package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
)

func HostList(c *gin.Context) {
	var hosts []v1.Hosts

	if err := db.Client().Find(&hosts).Error; err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "Database Error: %s", err.Error()), nil)
		return
	}

	hostList := &v1.HostsList{
		Hosts: hosts,
	}

	core.WriteResponse(c, nil, hostList)
}
