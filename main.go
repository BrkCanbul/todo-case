// @title Todo API
// @version 1.0
// @description Görev yönetimi servisi
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT formatında bearer token giriniz. Örn: "Bearer {token}"


package main


import (
	"todo-case/routes"
	"todo-case/services"
	"todo-case/controllers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "todo-case/docs"

)

func main() {
	r := gin.Default()

	todoService := services.NewMockTodoService()
	userService := services.NewMockUserService()

	todoController := controllers.NewTodoController(todoService)
	authController := controllers.NewAuthController(userService)

	routes.SetupRoutes(r, todoController, authController)
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
