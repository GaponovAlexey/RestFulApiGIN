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

func (s *TodoListsService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListsService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
func (s *TodoListsService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListsService) Update(userId, listId int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
