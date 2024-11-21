package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Basic CRUD operations for the User entity.

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"usuarios": Usuarios,
	})
}

func GetUserById(c *gin.Context) {
}

func CreateUser(c *gin.Context) {
}

func UpdateUser(c *gin.Context) {
}

func DeleteUser(c *gin.Context) {
}

// End of the basic CRUD operations for the User entity.
