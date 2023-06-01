package core

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"gmcm/pkg/log"
	"net/http"
)

type ErrReponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		log.Errorf("%#+v", err)
		coder := errors.ParseCoder(err)
		c.JSON(coder.HTTPStatus(), ErrReponse{
			Code:    coder.Code(),
			Message: coder.String(),
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
