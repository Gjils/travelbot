package repositories

import (
	"travel-bot/internal/entities"

	"github.com/google/uuid"
)

type UserSchema struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey"`
	TelegramId int
	Name string
	Age int
	Country string
	Town string
	Description string
}

func toSchema(user *entities.User) *UserSchema {
	return &UserSchema{
		Id: user.Id,
		TelegramId: user.TelegramId,
		Name: user.Name,
		Age: user.Age,
		Country: user.Location.Country,
		Town: user.Location.Town,
		Description: user.Description,
	}
}

func ToEntity(userSchema *UserSchema) *entities.User {
	location := entities.Location{
		Country: userSchema.Country,
		Town: userSchema.Town,
	}
	user := entities.User{
		Id: userSchema.Id,
		TelegramId: userSchema.TelegramId,
		Name: userSchema.Name,
		Age: userSchema.Age,
		Location: location,
		Description: userSchema.Description,
	}

	return &user
}