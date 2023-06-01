package routes

import (
	"encoding/base64"
	"fmt"
	v1 "gmcm/model/v1"
	"gmcm/pkg/db"
	"gmcm/pkg/log"
	"gmcm/pkg/middleware"
	"gmcm/pkg/middleware/auth"
	"gmcm/pkg/utils"
	"net/http"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	timeout           time.Duration = 24 * time.Hour
	maxRefresh        time.Duration = 24 * time.Hour
	APIServerIssuser                = "gmcm"
	APIServerAudience               = "macrowing.com"
)

type loginInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func newJWTAuth() middleware.AuthStrategy {
	ginjwt, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            utils.GetEnv("jwt.realm", ""),
		SigningAlgorithm: "HS256",
		Key:              []byte(utils.GetEnv("jwt.key", "")),
		Timeout:          timeout,
		MaxRefresh:       maxRefresh,
		Authenticator:    authenticator(),
		LoginResponse:    loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, nil)
		},
		RefreshResponse: refreshResponse(),
		PayloadFunc:     payloadFunc(),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return claims[jwt.IdentityKey]
		},
		IdentityKey:  middleware.UsernameKey,
		Authorizator: authorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return auth.NewJWTStrategy(*ginjwt)
}

func newAutoAuth() middleware.AuthStrategy {
	return auth.NewAutoStragegy(newBasicAuth().(auth.BasicStrategy), newJWTAuth().(auth.JWTStrategy))
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username string, password string) bool {
		user := &v1.User{}

		err := db.Client().Table("users").Where("name = ?", username).First(&user).Error
		if err != nil {
			return false
		}
		if err := user.Compare(password); err != nil {
			return false
		}

		return true
	})
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var login loginInfo
		var err error
		var user = &v1.User{}

		if c.Request.Header.Get("Authorization") != "" {
			login, err = parseWithHeader(c)
		} else {
			login, err = parseWithBody(c)
		}

		if err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		err2 := db.Client().Table("users").Where("name = ?", login.Username).First(&user).Error
		if err2 != nil {
			log.Errorf("get user information failed: %s", err2.Error())
			return "", jwt.ErrFailedAuthentication
		}

		if err = user.Compare(login.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		return user, nil
	}
}

func parseWithHeader(c *gin.Context) (loginInfo, error) {
	auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		log.Errorf("get basic string from Authorization header failed")
		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	payload, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		log.Errorf("decode basic string: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		log.Errorf("parse payload failed")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return loginInfo{
		Username: pair[0],
		Password: pair[1],
	}, nil
}

func parseWithBody(c *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Errorf("parse login paramaters: %s", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return login, nil
}

func refreshResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{
			"iss": APIServerIssuser,
			"aud": APIServerAudience,
		}

		if u, ok := data.(*v1.User); ok {
			claims[jwt.IdentityKey] = u.UserName
			claims["sub"] = u.UserName
			claims["id"] = u.ID
		}
		return claims
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(string); ok {
			log.Infof("user `%s` is authenticated", v)
			return true
		}
		return false
	}
}
