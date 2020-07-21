package controllers

import (
	"fmt"
	"net/http"

	"github.com/Kishan2595/Go-UserRegistrationLogin-API/models"
	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStatement := `INSERT INTO Users (Username, Password) VALUES ($1, $2) RETURNING id`

	id := 0
	row := models.DB.QueryRow(sqlStatement, input.Username, input.Password)
	err1 := row.Scan(&id)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	fmt.Println("New ID is:", id)
	c.JSON(http.StatusOK, gin.H{"data": input})
}
