package controllers

import (
	"fmt"
	"net/http"
	"strconv"
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
	userType := (c.GetString("user_type"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if userType == "user" {
		result := make([]models.ToDoList, 0)
		for _, todo := range *todos {
			if todo.UserId == int32(c.GetUint("user_id")) {
				result = append(result, todo)
			}
		}
		c.JSON(http.StatusOK, result)
		return
	}
	c.JSON(http.StatusOK, todos)
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

	fmt.Printf("AddToDoList %v", c.GetString("user_type"))
	fmt.Print(c.GetString("user_type"))
	userid := c.GetUint("user_id")

	if err := c.ShouldBindJSON(&todolist); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error()})
		return
	}
	todolist.UserId = int32(userid)

	created, err := ctl.srv.Create(&todolist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, created)
}

// AddToDoElement godoc
// @Summary      Yeni görev oleuştur
// @Description  Listeye Yeni bir görev  ekler
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        todoList  body      models.ToDo  true  "Yeni görev"
// @Success      200  {object}  models.ToDo
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /todos/elems [post]
func (ctl *TodoController) AddToDoElement(c *gin.Context) {
	var todoElement models.ToDo
	userid := int32(c.GetUint("user_id"))
	usertype := c.GetString("user_type")

	if err := c.ShouldBindJSON(&todoElement); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error()})
		return
	}

	created, err := ctl.srv.CreateElement(&todoElement, userid, usertype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, created)
}

// GetToDoElements godoc
// @Summary      görev listesi elemanlarını getir
// @Description  Listeye Yeni bir görev  ekler
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Success      200  {array}  models.ToDo
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /todos/elems [get]
func (ctl *TodoController) GetToDoElements(c *gin.Context) {
	usertype := c.GetString("user_type")
	userid := int32(c.GetUint("user_id"))
	todoElements, err := ctl.srv.GetAllElements(usertype, userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, todoElements)
}

// GetElementsByListId godoc
// @Summary      Belirli bir listeye ait görevleri getir
// @Description  Verilen liste kimliğine (id) göre görevleri döndürür
// @Tags         todos
// @Accept       json
// @Produce      json
// @Security BearerAuth
// @Param        id   query     int  true  "Liste Kimliği"
// @Success      200  {array}   models.ToDo
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /todos/elems [get]
func (ctl *TodoController) GetElementsByListId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Printf("GetElementsByListId %v", c.GetString("user_type"))

	userid := int32(c.GetUint("user_id"))
	usertype := c.GetString("user_type")

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "id parametresi geçersiz"})
		return
	}

	todoElements, err := ctl.srv.GetElementsByListId(uint(id), userid, usertype)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, todoElements)
}
