package main

import (
	"github.com/Kishan2595/Go-UserRegistrationLogin-API/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.main()

	r.POST("r/egister", controllers.CreateUser)
	r.POST("r/login", controllers.CreateUser)

	r.Run()
}
