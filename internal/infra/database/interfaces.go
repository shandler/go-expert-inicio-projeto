package database

import "github.com/shandler/go-expert-inicio-projeto/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emaild string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) (*entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Updete(product *entity.Product) error
	Delete(id string) error
}
