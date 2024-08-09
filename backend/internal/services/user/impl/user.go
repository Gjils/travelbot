package service

import (
	"travel-bot/internal/services/user/api"

	"github.com/google/uuid"
)

type UserService struct {
	rep UserRepository
}

func GetUserService(rep UserRepository) api.UserService {
	return &UserService{
		rep: rep,
	}
}

func (serv *UserService) AddUser(userInfo api.UserDTO) (uuid.UUID, error) {
	id := uuid.New()
	user, err := ToEntity(userInfo, id)
	if err != nil {
		return uuid.UUID{}, err
	}

	if err := serv.rep.AddUser(user); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (serv *UserService) UpdateUser(id uuid.UUID, userInfo api.UserDTO) error {
	user, err := ToEntity(userInfo, id)
	if err != nil {
		return err
	}

	if err := serv.rep.UpdateUser(id, user); err != nil {
		return err
	}

	return nil
}

func (serv *UserService) RemoveUser(id uuid.UUID) error {
	if err := serv.rep.RemoveUser(id); err != nil {
		return err
	}
	
	return nil
}

func (serv *UserService) GetUserById(id uuid.UUID) (*api.UserDTO, error) {
	user, err := serv.rep.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userDTO, err := ToDTO(user)
	if err != nil {
		return nil, err
	}

	return &userDTO, nil
}