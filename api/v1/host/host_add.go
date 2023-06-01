package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
	"gmcm/pkg/log"
	"strings"
	"time"
)

func HostAdd(c *gin.Context) {
	var host *v1.Hosts
	err := c.ShouldBindJSON(&host)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "用户传参错误: %s. 400 Bad Request", err.Error()), nil)
		return
	}

	host.ObjectMeta.SetCreateAt(time.Now())
	host.ObjectMeta.SetUpdateAt(time.Now())

	clone := host
	ips := strings.Split(host.IP, ",")

	for _, ip := range ips {
		_, err = newSSHClient(*clone, ip)
		if err != nil {
			log.Errorf("无法与目标主机 %s 建立链接", ip)
			continue
		}
	}

	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrHostConnectionByPass, "无法与目标主机建立链接", err.Error()), nil)
		return
	}

	for _, ip := range ips {
		clone.IP = ip
		if len(ips) > 1 {
			clone.ID = clone.ID + 1
		}
		err = db.Client().Table("hosts").Create(&clone).Error
	}

	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "数据库错误. %s", err.Error()), nil)
		return
	}

	err = refreshHostTemplate()
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, "刷新模板错误. %s", err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"successes": "ok"})
}
