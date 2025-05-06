package controllers

import (
	"net/http"
	"todo-case/models"
	"todo-case/services"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	srv services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{srv: service}

}
func (ctl *TodoController) GetLists(c *gin.Context) {
	todos := ctl.srv.GetAll()
	c.JSON(http.StatusOK, todos)

}

func AddTodoCase(c *gin.Context) {

	c.JSON(http.StatusNotImplemented, gin.H{"error": "this method is not implemented"})
}
func RemoveTodoCase(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "this method is not implemented"})
}
func (ctl *TodoController) AddToDoList(c *gin.Context) {
	var todolist models.ToDoList
	if err := c.ShouldBindJSON(&todolist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	created := ctl.srv.Create(&todolist)

	c.JSON(http.StatusOK, created)
}

func RemoveFromList(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "this method is not implemented"})
}
func EditTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "this method is not implemented"})

}
func EditList(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"error": "this method is not implemented"})

}
