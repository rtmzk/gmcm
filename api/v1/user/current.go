package user

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
)

func Current(c *gin.Context) {
	var u = &v1.User{}
	err := db.Client().Table("users").Where("name = ?", c.GetString("username")).First(&u).Error
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, "Could not get current user infomation in database."), nil)
		return
	}
	core.WriteResponse(c, nil, u)
}
