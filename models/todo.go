package models

import (
	"time"
)


type ToDoList struct {
	ListId uint `json:"list_id" example:"1"`
	ListName string `json:"list_name" example:"list name"`
	CreateDate time.Time `json:"create_date" example:"2023-10-01T12:00:00Z" swaggerignore:"true"`
	UpdateDate time.Time `json:"update_date" example:"2023-10-01T12:00:00Z" swaggerignore:"true"`
	RemoveDate time.Time `json:"remove_date" example:"2023-10-01T12:00:00Z" swaggerignore:"true"`
	CompleteStatus float32 `json:"complete_status" example:"0.5"`
	UserId int32 `json:"user_id" example:"1"`

}
type ToDo struct {
	TodoId uint `json:"todo_id" example:"1"`
	TodolistId uint `json:"todolist_id" example:"1"`
	CreateDate time.Time `json:"create_date" example:"2023-10-01T12:00:00Z" swaggerignore:"true"`
	UpdateDate time.Time `json:"update_date" example:"2023-10-01T12:00:00Z" swaggerignore:"true"`
	RemoveDate time.Time `json:"remove_date" example:"2023-10-01T12:00:00Z" swaggerignore:"true"`
	Content string `json:"content" example:"todo content"`
	IsCompleted bool `json:"is_completed" example:"false"`
}
