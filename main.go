package main

import (
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/controllers"
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/models"
	"github.com/gin-gonic/gin"
)

func Connectdb() {
	r := gin.Default()

	models.Connectdb()

	r.POST("r/egister", controllers.CreateUser)
	r.POST("r/login", controllers.CreateUser)

	r.Run()
}
