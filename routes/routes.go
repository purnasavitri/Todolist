package routes

import (
	"todolist/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/user", controller.GetAllUser)
	r.POST("/user", controller.CreateUser)
	r.GET("/user/:id", controller.GetUserByID)
	r.PUT("/user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
}