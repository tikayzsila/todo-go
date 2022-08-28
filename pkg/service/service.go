package service

import (
	restgo "todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user restgo.User) (int, error)
}

type TodoList interface{}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}

}
