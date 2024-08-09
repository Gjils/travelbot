package entities

import "github.com/google/uuid"

type User struct {
	Id uuid.UUID
	TelegramId int
	Name string
	Age int
	Location Location
	Description string
}