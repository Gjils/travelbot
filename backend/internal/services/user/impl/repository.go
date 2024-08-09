package service

import (
	"travel-bot/internal/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	AddUser(userInfo entities.User) error
	UpdateUser(id uuid.UUID, userInfo entities.User) error
	RemoveUser(id uuid.UUID) error
	GetUserById(id uuid.UUID) (entities.User, error)
}