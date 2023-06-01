package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
)

func GetInitStatus(c *gin.Context) {
	var status *v1.Status

	err := db.Client().Table("statuses").Find(&status).Error
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "err: %s", err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, &status)
}
