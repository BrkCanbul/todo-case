package routes

import(
	"github.com/gin-gonic/gin"
	"todo-case/controllers"
)
func Setup(router *gin.Engine){
	router.POST("/login",controllers.Login)
	
}