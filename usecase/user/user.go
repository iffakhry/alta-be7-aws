package user

import (
	_entities "be7/layered/entities"
	_userRepository "be7/layered/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.userRepository.GetAll()
	return users, err
}

func (uuc *UserUseCase) GetById(id int) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.GetById(id)
	return user, rows, err
}

func (uuc *UserUseCase) Insert(user _entities.User) (_entities.User, int, error) {
	user, rows, err := uuc.userRepository.Insert(user)
	return user, rows, err
}
