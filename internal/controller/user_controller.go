package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	// Implement your logic to get user information
	ctx.JSON(200, gin.H{
		"message": "User information",
	})
}
