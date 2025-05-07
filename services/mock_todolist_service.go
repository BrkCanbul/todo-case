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

	DeleteList(id uint, userId int32, userType string) error
	DeleteTodo(id uint, userId int32, userType string) error

	UpdateList(updatedList *models.ToDoList, userId int32, userType string) (*models.ToDoList, error)
	UpdateTodo(updatedTodo *models.ToDo, userId int32, userType string) (*models.ToDo, error)
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
func (s *MockTodoService) DeleteList(id uint, userId int32, userType string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, list := range s.todoLists {
		if list.ListId == id {
			if userType == "admin" || list.UserId == userId {
				s.todoLists = append(s.todoLists[:i], s.todoLists[i+1:]...)
				// O listeye ait görevleri de silelim
				filtered := make([]models.ToDo, 0)
				for _, todo := range s.todos {
					if todo.TodolistId != id {
						filtered = append(filtered, todo)
					}
				}
				s.todos = filtered
				return nil
			}
			return errors.New("bu listeyi silmeye yetkiniz yok")
		}
	}
	return errors.New("liste bulunamadı")
}

func (s *MockTodoService) UpdateList(updatedList *models.ToDoList, userId int32, userType string) (*models.ToDoList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, list := range s.todoLists {
		if list.ListId == updatedList.ListId {
			if userType == "admin" || list.UserId == userId {
				updatedList.CreateDate = list.CreateDate
				updatedList.RemoveDate = list.RemoveDate
				updatedList.UpdateDate = time.Now()
				s.todoLists[i] = *updatedList
				return updatedList, nil
			}
			return nil, errors.New("bu listeyi güncellemeye yetkiniz yok")
		}
	}
	return nil, errors.New("liste bulunamadı")
}

func (s *MockTodoService) DeleteTodo(id uint, userId int32, userType string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.TodoId == id {
			listOwnerId := s.findListOwner(todo.TodolistId)
			if userType == "admin" || listOwnerId == userId {
				s.todos = append(s.todos[:i], s.todos[i+1:]...)
				return nil
			}
			return errors.New("bu görevi silmeye yetkiniz yok")
		}
	}
	return errors.New("görev bulunamadı")
}

func (s *MockTodoService) UpdateTodo(updatedTodo *models.ToDo, userId int32, userType string) (*models.ToDo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, todo := range s.todos {
		if todo.TodoId == updatedTodo.TodoId {
			listOwnerId := s.findListOwner(todo.TodolistId)
			if userType == "admin" || listOwnerId == userId {
				updatedTodo.CreateDate = todo.CreateDate
				updatedTodo.RemoveDate = todo.RemoveDate
				updatedTodo.UpdateDate = time.Now()
				s.todos[i] = *updatedTodo
				return updatedTodo, nil
			}
			return nil, errors.New("bu görevi güncellemeye yetkiniz yok")
		}
	}
	return nil, errors.New("görev bulunamadı")
}

// Yardımcı fonksiyon
func (s *MockTodoService) findListOwner(listId uint) int32 {
	for _, list := range s.todoLists {
		if list.ListId == listId {
			return list.UserId
		}
	}
	return -1
}
