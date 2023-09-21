package database

import "github.com/shandler/go-expert-inicio-projeto/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}
