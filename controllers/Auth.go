package controllers

import (
	"net/http"

	"github.com/Kishan2595/Go-UserRegistrationLogin-API/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginUserInput struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO Users (Username, Password) VALUES ($1, $2)`

	_, err1 := models.DB.Query(sqlStatement, input.Username, input.Password)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func LoginUser(c *gin.Context) {
	var input LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlStatement := `Select * From Users where Username = ($1) and Password = ($2)`
	row := models.DB.QueryRow(sqlStatement, input.Username, input.Password)
	err := row.Scan(&input.Username, &input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": input})

}

func RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
