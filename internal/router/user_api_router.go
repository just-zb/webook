package router

import (
	"github.com/gin-gonic/gin"
	"github.com/just-zb/webook/internal/controller"
)

type UserAPIRouter struct {
	userController *controller.UserController
}

func NewUserAPIRouter(userController *controller.UserController) *UserAPIRouter {
	return &UserAPIRouter{
		userController: userController,
	}
}
func (router *UserAPIRouter) Register(r *gin.RouterGroup) {
	r.GET("/", router.userController.GetUser)
}
