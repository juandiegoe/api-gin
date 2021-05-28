package routes

import (
	"github.com/JuanDiegoE/api-gin/middlewares"

	"github.com/gin-gonic/gin"
)


func Setup(server *gin.Engine) {

	server.GET("/", middlewares.Saludo)
	server.POST("/user", middlewares.CreateUser)
	server.POST("/login", middlewares.LoginUser)
	server.GET("/logged", middlewares.UserLogged)
	server.GET("/logout", middlewares.Logout)
}
