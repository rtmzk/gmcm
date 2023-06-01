package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/middleware"
	"strings"
)

type AutoStrategy struct {
	basic BasicStrategy
	jwt   JWTStrategy
}

const authHeaderCount = 2

var _ middleware.AuthStrategy = &AutoStrategy{}

func NewAutoStragegy(basic BasicStrategy, jwt JWTStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := middleware.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeaderCount {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong"),
				nil,
			)

			c.Abort()
			return
		}
		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "Bearer":
			operator.SetStrategy(a.jwt)
		default:
			core.WriteResponse(c, errors.WithCode(code.ErrSignatureInvalid, "unrecongnized Authorization header"), nil)

			c.Abort()
			return
		}

		operator.AuthFunc()(c)
		c.Next()
	}
}
