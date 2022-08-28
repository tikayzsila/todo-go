package repository

import (
	restgo "todo-app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user restgo.User) (int, error)
	GetUser(username, password string) (restgo.User, error)
}

type TodoList interface {
	Create(userId int, list restgo.TodoList) (int, error)
}

type TodoItem interface{}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}

}
