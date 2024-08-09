package service

import (
	"travel-bot/internal/entities"
	"travel-bot/internal/services/user/api"

	"github.com/google/uuid"
)


func ToEntity(dto api.UserDTO, id uuid.UUID) (entities.User, error) {
	user := entities.User{
		TelegramId: dto.TelegramId,
		Name: dto.Name,
		Age: dto.Age,
		Location: entities.Location{
			Country: dto.Country,
			Town: dto.Town,
		},
		Description: dto.Description,
	}

	return user, nil
}

func ToDTO(user entities.User) (api.UserDTO, error) {
	userDTO := api.UserDTO{
		TelegramId: user.TelegramId,
		Name: user.Name,
		Age: user.Age,
		Country: user.Location.Country,
		Town: user.Location.Town,
		Description: user.Description,
	}

	return userDTO, nil
}