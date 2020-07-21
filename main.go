package main

import (
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/controllers"
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Connectdb()

	r.POST("/register", controllers.CreateUser)
	//r.POST("/login", controllers.LoginUser)

	r.Run()
}
