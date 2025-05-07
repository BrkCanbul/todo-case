package routes

import (
	"todo-case/controllers"
	"todo-case/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, todoController *controllers.TodoController, authController *controllers.AuthController) {
	r.POST("/login", authController.Login)

	authGroup := r.Group("/todos")
	authGroup.Use(utils.AuthMiddleware())
	{
		authGroup.GET("/", todoController.GetLists)
		authGroup.POST("/", todoController.AddToDoList)
		authGroup.GET("/elems", todoController.GetToDoElements)
		authGroup.GET("/elems/:id", todoController.GetElementsByListId)
		authGroup.POST("/elems", todoController.AddToDoElement)
	}
}
