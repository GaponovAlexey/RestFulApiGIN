package service

import (
	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/repository"
)

type TodoListsService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListsService {
	return &TodoListsService{repo: repo}
}

func (s *TodoListsService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
