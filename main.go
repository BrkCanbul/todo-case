package main

import (
	"todo-case/routes"
	"todo-case/services"
	"todo-case/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	todoService := services.NewMockTodoService()
	userService := services.NewMockUserService()

	todoController := controllers.NewTodoController(todoService)
	authController := controllers.NewAuthController(userService)

	routes.SetupRoutes(r, todoController, authController)

	r.Run(":8080")
}
