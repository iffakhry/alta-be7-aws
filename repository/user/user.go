package user

import (
	_entities "be7/layered/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAll() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (ur *UserRepository) GetById(id int) (_entities.User, int, error) {
	var user _entities.User
	tx := ur.database.Find(&user, id)
	if tx.Error != nil {
		return user, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return user, 0, nil
	}
	return user, int(tx.RowsAffected), nil
}

func (ur *UserRepository) Insert(user _entities.User) (_entities.User, int, error) {

	tx := ur.database.Create(&user)
	if tx.Error != nil {
		return _entities.User{}, 0, tx.Error
	}
	return user, 1, nil

}
