package entities

import (
	"time"

	"github.com/google/uuid"
)

type Trip struct {
	Id uuid.UUID
	OwnerId int
	Name string
	Description string
	Locations []TripLocation
	Participators []int
}

type TripLocation struct {
	Location
	Id uuid.UUID
	StartDate time.Time
	EndDate time.Time
}