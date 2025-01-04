package routes

import (
	controller "github.com/bhargavduddugunta/golang-jwt-project/controllers"
	"github.com/bhargavduddugunta/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.UserAuthenticate())
	incomingRoutes.POST("/userdata",controller.GetUsers())
	incomingRoutes.POST("/users/:user_id",controller.GetUser())
}