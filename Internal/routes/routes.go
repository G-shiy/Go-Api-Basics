package routes

import (
	"FintechPrototype/Internal/User"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user", User.GetUsers)
	r.GET("/user/:id", User.GetUserById)
	r.POST("/user", User.CreateUser)
	r.PUT("/user/:id", User.UpdateUser)
	r.DELETE("/user/:id", User.DeleteUser)
}
