package model

import "github.com/google/uuid"

type Profiles struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	User      User
}
