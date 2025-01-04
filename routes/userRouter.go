package routes

import (
	controller "github.com/bhargavduddugunta/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func UserAuthenticate(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup",controller.Signup())
	incomingRoutes.POST("users/login",controller.Login())
}