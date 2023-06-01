package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	v1 "gmcm/api/v1"
	"gmcm/api/v1/host"
	"gmcm/api/v1/user"
	"gmcm/pkg/middleware"
	"gmcm/pkg/middleware/auth"
	_ "gmcm/statik"
)

func SetRoutes(r *gin.Engine) {
	var statikFS static.ServeFileSystem
	statikFS = &GinFS{}
	statikFS.(*GinFS).FS, _ = fs.New()

	r.Use(middleware.Cors(), middleware.Logger(), gin.Recovery(), static.Serve("/", statikFS))

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)

	r.POST("/api/login", jwtStrategy.LoginHandler)
	r.POST("/api/logout", jwtStrategy.LogoutHandler)
	r.POST("/api/refresh", jwtStrategy.RefreshHandler)

	r.GET("/", v1.UI)

	r.StaticFS("/ui", statikFS)

	auto := newAutoAuth()

	r.GET("/healthz", v1.GetHealthz)

	api := r.Group("/api")
	{
		u := api.Group("/users")
		{
			u.Use(auto.AuthFunc())
			u.GET("/current", user.Current)
		}

		h := api.Group("/host")
		{
			h.Use(auto.AuthFunc())
			h.POST("/check", host.ConnectionCheck)
			h.POST("/add", host.HostAdd)
			h.GET("/list", host.HostList)
			h.DELETE("/delete/:id", host.HostDelete)
			h.GET("/devices", host.GetHostDevices)
			h.POST("/init", host.HostInit)
			h.GET("/init/log/:offset", host.GetInstallLogs)
			h.GET("/envc/prepare", host.EnvCheckPrepare)
			h.POST("/envc", host.EnvcAction)
			h.GET("/init/status", host.GetInitStatus)
			h.POST("/init/status/update", host.SetInitStatus)
		}
	}
}
