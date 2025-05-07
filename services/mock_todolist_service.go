package services

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"todo-case/models"
)

type TodoService interface {
	GetAll() (*[]models.ToDoList, error)
	GetAllElements(usertype string, userid int32) (*[]models.ToDo, error)
	GetElementsByListId(id uint, userid int32, usertype string) (*[]models.ToDo, error)
	Create(todo *models.ToDoList) (*models.ToDoList, error)
	CreateElement(todoElement *models.ToDo, userid int32, usertype string) (*models.ToDoList, error)
}

type MockTodoService struct {
	todoLists []models.ToDoList
	todos     []models.ToDo
	mu        sync.Mutex
	listId    uint
	todoId    uint
}

func NewMockTodoService() *MockTodoService {
	return &MockTodoService{
		todoLists: []models.ToDoList{},
		todos:     []models.ToDo{},
		listId:    1,
		todoId:    1,
	}
}

func (s *MockTodoService) GetAll() (*[]models.ToDoList, error) {
	s.mu.Lock()
	result := make([]models.ToDoList, len(s.todoLists))
	s.mu.Unlock()
	for i := range s.todoLists {
		print(s.todoLists[i].ListId)
		result[i] = models.ToDoList{
			ListId:         s.todoLists[i].ListId,
			ListName:       s.todoLists[i].ListName,
			CreateDate:     s.todoLists[i].CreateDate,
			RemoveDate:     s.todoLists[i].RemoveDate,
			CompleteStatus: s.todoLists[i].CompleteStatus,
			UserId:         s.todoLists[i].UserId,
		}
	}
	return &result, nil
}

func (s *MockTodoService) GetUserListIds(id int32) []int {
	usertodoids := make([]int, 0)
	for _, elem := range s.todoLists {
		if elem.UserId == id {
			usertodoids = append(usertodoids, int(elem.ListId))
		}
	}
	return usertodoids
}

func (s *MockTodoService) GetElementsByListId(id uint, userid int32, usertype string) (*[]models.ToDo, error) {
	s.mu.Lock()
	result := make([]models.ToDo, 0)
	s.mu.Unlock()
	usertodoids := s.GetUserListIds(userid)
	if usertype == "user" {
		for _, elem := range usertodoids {
			for _, todo := range s.todos {
				if todo.TodolistId == uint(elem) && todo.TodolistId == id {
					result = append(result, models.ToDo{
						TodoId:      todo.TodoId,
						TodolistId:  todo.TodolistId,
						CreateDate:  todo.CreateDate,
						RemoveDate:  todo.RemoveDate,
						Content:     todo.Content,
						IsCompleted: todo.IsCompleted,
					})
				}
			}
		}
	} else {
		for _, todo := range s.todos {
			if todo.TodolistId == id {
				result = append(result, models.ToDo{
					TodoId:      todo.TodoId,
					TodolistId:  todo.TodolistId,
					CreateDate:  todo.CreateDate,
					RemoveDate:  todo.RemoveDate,
					Content:     todo.Content,
					IsCompleted: todo.IsCompleted,
				})
			}
		}
	}

	if len(result) == 0 {
		return nil, errors.New(fmt.Sprintf("id with %d couldn't found in list", id))
	}
	return &result, nil
}

func (s *MockTodoService) GetAllElements(usertype string, userid int32) (*[]models.ToDo, error) {
	s.mu.Lock()
	result := make([]models.ToDo, len(s.todos))
	s.mu.Unlock()
	usertodoids := s.GetUserListIds(userid)
	if usertype == "user" {
		for _, elem := range usertodoids {
			for _, todo := range s.todos {
				if todo.TodolistId == uint(elem) {
					result = append(result, models.ToDo{
						TodoId:      todo.TodoId,
						TodolistId:  todo.TodolistId,
						CreateDate:  todo.CreateDate,
						RemoveDate:  todo.RemoveDate,
						Content:     todo.Content,
						IsCompleted: todo.IsCompleted,
					})
				}
			}
		}
	} else {
		for i := range s.todos {
			result[i] = models.ToDo{
				TodoId:      s.todos[i].TodoId,
				TodolistId:  s.todos[i].TodolistId,
				CreateDate:  s.todos[i].CreateDate,
				RemoveDate:  s.todos[i].RemoveDate,
				Content:     s.todos[i].Content,
				IsCompleted: s.todos[i].IsCompleted,
			}
		}
	}
	return &result, nil
}

func (s *MockTodoService) Create(list *models.ToDoList) (*models.ToDoList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	list.ListId = s.listId

	s.listId++

	s.todoLists = append(s.todoLists, models.ToDoList{
		ListId:         list.ListId,
		ListName:       list.ListName,
		CreateDate:     time.Now(),
		RemoveDate:     time.Now().Add(time.Hour * 24 * 30 * 12),
		UpdateDate:     time.Now(),
		CompleteStatus: list.CompleteStatus,
		UserId:         list.UserId,
	})
	return list, nil
}

func (s *MockTodoService) CreateElement(todoElement *models.ToDo, userid int32, usertype string) (*models.ToDoList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var selected *models.ToDoList = nil
	userlistids := s.GetUserListIds(userid)
	for _, elem := range s.todoLists {
		if elem.ListId == todoElement.TodolistId {
			selected = &elem
		}
	}
	if selected == nil {
		return nil, errors.New(fmt.Sprintf("id with %d couldn't found in list", todoElement.TodolistId))
	}

	todoElement.TodoId = s.todoId
	s.todoId++
	todoElement.CreateDate = time.Now()
	todoElement.RemoveDate = time.Now().Add(time.Hour * 24 * 30 * 12)
	todoElement.UpdateDate = time.Now()

	todoElement.IsCompleted = false
	todoElement.TodolistId = selected.ListId
	todoElement.Content = todoElement.Content
	isInList := false
	for _, elem := range userlistids {
		if elem == int(selected.ListId) {
			isInList = true
		}
	}
	if !isInList && usertype == "user" {
		return nil, errors.New(fmt.Sprintf("id with %d couldn't found in list", todoElement.TodolistId))
	}
	s.todos = append(s.todos, *todoElement)
	return selected, nil
}
