package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
)

func SetInitStatus(c *gin.Context) {
	var status *v1.Status
	err := c.ShouldBindJSON(&status)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Bad Request"), nil)
		return
	}

	err = db.Client().Table("statuses").Where("id = ?", 1).Update("status", &status.InitStatus).Error
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "Update init status failed. err: %s", err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, map[string]string{"successes": "ok"})
}
