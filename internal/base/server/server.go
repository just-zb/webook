package server

import (
	"github.com/gin-gonic/gin"
	"github.com/just-zb/webook/internal/router"
)

func newHttpServer(debug bool,
	userApiRouter *router.UserAPIRouter) *gin.Engine {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.GET("/healthz", func(ctx *gin.Context) { ctx.String(200, "OK") })
	userGroup := r.Group("/user")
	userApiRouter.Register(userGroup)
	return r
}
