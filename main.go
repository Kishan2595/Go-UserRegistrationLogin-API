package main

import (
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/controllers"
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/")

	models.ConnectDatabase()

	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/register", controllers.RegisterForm)
	r.GET("/login", controllers.LoginForm)
	r.Run()
}
