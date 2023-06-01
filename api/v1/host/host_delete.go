package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
	"gmcm/pkg/log"
)

func HostDelete(c *gin.Context) {
	log.Debug("Host delete function called.")

	hostId := c.Param("id")
	err := db.Client().Table("hosts").Where("id = ?", hostId).Unscoped().Delete(v1.Hosts{}).Error
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "主机删除失败: %s", err.Error()), nil)
		return
	}
	err = refreshHostTemplate()
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, "刷新模板错误. %s", err.Error()), nil)
		return
	}
	core.WriteResponse(c, nil, map[string]string{"successes": "true"})
}
