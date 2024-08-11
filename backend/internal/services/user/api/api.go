package api

import (
	"github.com/google/uuid"
)

type UserService interface {
	AddUser(userInfo *UserDTO) (uuid.UUID, error)
	UpdateUser(id uuid.UUID, userInfo *UserDTO) error
	RemoveUser(id uuid.UUID) error
	GetUserById(id uuid.UUID) (*UserDTO, error)
}