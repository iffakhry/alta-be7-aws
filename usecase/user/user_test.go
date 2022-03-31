package user

import (
	_entities "be7/layered/entities"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// jangan lupa download testify
// go get -u github.com/stretchr/testify

func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, err := userUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "alta", data[0].Name)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, err := userUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}

func TestGetById(t *testing.T) {
	t.Run("TestGetByIdSuccess", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepository{})
		data, rows, err := userUseCase.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, "alta", data.Name)
		assert.Equal(t, 1, rows)
	})

	t.Run("TestGetByIdError", func(t *testing.T) {
		userUseCase := NewUserUseCase(mockUserRepositoryError{})
		data, rows, err := userUseCase.GetById(1)
		assert.NotNil(t, err)
		assert.Equal(t, 0, rows)
		assert.Equal(t, _entities.User{}, data)
	})
}

// === mock success ===
type mockUserRepository struct{}

func (m mockUserRepository) GetAll() ([]_entities.User, error) {
	return []_entities.User{
		{Name: "alta", Email: "alta@mail.com", Password: "12345"},
	}, nil
}

func (m mockUserRepository) GetById(id int) (_entities.User, int, error) {
	return _entities.User{
		Name: "alta", Email: "alta@mail.com", Password: "12345",
	}, 1, nil
}

func (m mockUserRepository) Insert(data _entities.User) (_entities.User, int, error) {
	return _entities.User{
		Name: "alta", Email: "alta@mail.com", Password: "12345",
	}, 1, nil
}

// === mock error ===

type mockUserRepositoryError struct{}

func (m mockUserRepositoryError) GetAll() ([]_entities.User, error) {
	return nil, fmt.Errorf("error")
}

func (m mockUserRepositoryError) GetById(id int) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error get data user")
}

func (m mockUserRepositoryError) Insert(data _entities.User) (_entities.User, int, error) {
	return _entities.User{}, 0, fmt.Errorf("error insert data user")
}
