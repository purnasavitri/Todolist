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

	r.GET("/category", controller.GetAllCategory)
	r.GET("/category/:id_category", controller.GetCategoryByID)
	r.POST("/category", controller.CreateCategory)
	r.PUT("/category/:id_category", controller.UpdateCategory)
	r.DELETE("/category/:id_category", controller.DeleteCategory)
}
