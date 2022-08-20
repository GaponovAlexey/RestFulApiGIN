package service

import (
	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}
func (s *TodoItemService) GetById(userId, ItemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, ItemId)
}
func (s *TodoItemService) Delete(userId, ItemId int) error {
	return s.repo.Delete(userId, ItemId)
}
func (s *TodoItemService) Update(userId, ItemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, ItemId, input)
}
