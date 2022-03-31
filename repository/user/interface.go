package user

import (
	_entities "be7/layered/entities"
)

type UserRepositoryInterface interface {
	GetAll() ([]_entities.User, error)
	GetById(id int) (_entities.User, int, error)
	Insert(data _entities.User) (_entities.User, int, error)
}
