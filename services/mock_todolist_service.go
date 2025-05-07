package services

import (
	"errors"
	"fmt"
	"sync"
	"todo-case/models"
	"time"
)

type TodoService interface {
	GetAll() (*[]models.ToDoList, error)
	Create(todo *models.ToDoList) (*models.ToDoList, error)
	CreateElement(todoElement *models.ToDo) (*models.ToDoList, error)
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
			ListId:        s.todoLists[i].ListId,
			ListName:      s.todoLists[i].ListName,
			CreateDate:    s.todoLists[i].CreateDate,
			RemoveDate:    s.todoLists[i].RemoveDate,
			CompleteStatus:s.todoLists[i].CompleteStatus,
			UserId:        s.todoLists[i].UserId,
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
		CompleteStatus: list.CompleteStatus,
		UserId:         list.UserId,
	})
	return list, nil
}

func (s *MockTodoService) CreateElement(todoElement *models.ToDo) (*models.ToDoList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var selected *models.ToDoList = nil

	for _, elem := range s.todoLists {
		if elem.ListId == todoElement.TodolistId {
			selected = &elem
		}
	}
	if selected == nil {
		return nil, errors.New(fmt.Sprintf("id with %d couldn't found in list",todoElement.TodolistId))
	}

	todoElement.TodoId = s.todoId
	s.todoId++
	s.todos = append(s.todos, *todoElement)

	return selected, nil
}
