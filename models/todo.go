package models

import (
	"google.golang.org/genproto/googleapis/type/datetime"
)

type ToDoList struct {
	listId uint
	listName string
	createDate datetime.DateTime
	removeDate datetime.DateTime
	completeStatus float32

}
type ToDo struct {
	todoId uint
	todolistId uint
	createDate datetime.DateTime
	removeDate datetime.DateTime
	content string
	isCompleted bool
}
