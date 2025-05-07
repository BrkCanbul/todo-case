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

// GetLists godoc
// @Summary      Tüm görevleri getir
// @Description  Kullanıcının tüm görevlerini listeler
// @Tags         todos
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.ToDo
// @Failure      500  {object}  models.ErrorResponse
// @Router       /todos [get]
func (ctl *TodoController) GetLists(c *gin.Context) {
	todos, err := ctl.srv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// AddTodoCase godoc
// @Summary      Yeni görev ekle
// @Description  Bu method henüz implemente edilmemiştir.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Router       /todos [post]
func AddTodoCase(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.ErrorResponse{Error: "this method is not implemented"})
}

// RemoveTodoCase godoc
// @Summary      Görev sil
// @Description  Bu method henüz implemente edilmemiştir.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Router       /todos/{id} [delete]
func RemoveTodoCase(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.ErrorResponse{Error: "this method is not implemented"})
}

// AddToDoList godoc
// @Summary      Yeni görev listesi oluştur
// @Description  Yeni bir görev listesi ekler
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        todoList  body      models.ToDoList  true  "Yeni görev listesi"
// @Success      200  {object}  models.ToDoList
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /todos [post]
func (ctl *TodoController) AddToDoList(c *gin.Context) {
	var todolist models.ToDoList
	if err := c.ShouldBindJSON(&todolist); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error()})
		return
	}


	created, err := ctl.srv.Create(&todolist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, created)
}

// RemoveFromList godoc
// @Summary      Listeyi sil
// @Description  Bu method henüz implemente edilmemiştir.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Router       /todos/{id} [delete]
func RemoveFromList(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.ErrorResponse{Error: "this method is not implemented"})
}

// EditTodo godoc
// @Summary      Görev düzenle
// @Description  Bu method henüz implemente edilmemiştir.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Router       /todos/{id} [put]
func EditTodo(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.ErrorResponse{Error: "this method is not implemented"})
}

// EditList godoc
// @Summary      Listeyi düzenle
// @Description  Bu method henüz implemente edilmemiştir.
// @Tags         todos
// @Accept       json
// @Produce      json
// @Router       /todos/{id} [put]
func EditList(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, models.ErrorResponse{Error: "this method is not implemented"})
}
