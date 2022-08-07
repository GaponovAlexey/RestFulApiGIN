package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/gaponovalexey/todo-app"
	"github.com/gaponovalexey/todo-app/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

const salt = "d1dadae255gg211"

func  generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
