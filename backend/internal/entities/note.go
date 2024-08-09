package entities

import "github.com/google/uuid"

type Note struct {
	Id uuid.UUID
	Title string
	Text string
	Images []Image
	Files []File
	IsPrivate bool
}

type Image struct {
	Id uuid.UUID
}

type File struct {
	Id uuid.UUID
}