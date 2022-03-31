package user

import (
	_entities "be7/layered/entities"
)

type UserUseCaseInterface interface {
	GetAll() ([]_entities.User, error)
	GetById(id int) (_entities.User, int, error)
	Insert(user _entities.User) (_entities.User, int, error)
	// Update(id int, user _entities.User) (_entities.User, error)
}
