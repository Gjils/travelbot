package service

import (
	"travel-bot/internal/services/user/api"

	"github.com/google/uuid"
)

type userService struct {
	rep UserRepository
}

func GetUserService(rep UserRepository) api.UserService {
	return &userService{
		rep: rep,
	}
}

func (serv *userService) AddUser(userInfo *api.UserDTO) (uuid.UUID, error) {
	id := uuid.New()
	user := ToEntity(userInfo, id)

	if err := serv.rep.AddUser(user); err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

func (serv *userService) UpdateUser(id uuid.UUID, userInfo *api.UserDTO) error {
	user := ToEntity(userInfo, id)

	if err := serv.rep.UpdateUser(id, user); err != nil {
		return err
	}

	return nil
}

func (serv *userService) RemoveUser(id uuid.UUID) error {
	if err := serv.rep.RemoveUser(id); err != nil {
		return err
	}
	
	return nil
}

func (serv *userService) GetUserById(id uuid.UUID) (*api.UserDTO, error) {
	user, err := serv.rep.GetUserById(id)
	if err != nil {
		return nil, err
	}

	userDTO := ToDTO(user)

	return userDTO, nil
}