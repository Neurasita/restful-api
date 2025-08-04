package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Gender rune

func GenderFromString(g string) (Gender, error) {
	switch g {
	case "m", "M":
		return GenderMale, nil
	case "f", "F":
		return GenderFemale, nil
	default:
		return GenderNil, errors.New("invalid user gender")
	}
}

func GenderFromRune(g rune) (Gender, error) {
	switch g {
	case 'm', 'M':
		return GenderMale, nil
	case 'f', 'F':
		return GenderFemale, nil
	default:
		return GenderNil, errors.New("invalid user gender")
	}
}

// User gender
const (
	GenderMale   Gender = 'm'
	GenderFemale Gender = 'f'
	GenderNil    Gender = ' '
)

// User Profile Table
type Profile struct {
	// PK
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	// Main data
	FirstName    *string    `gorm:"type:varchar(64)"`
	LastName     *string    `gorm:"type:varchar(64)"`
	PhotoURL     *string    `gorm:"type:varchar(255);uniqueIndex:uidx_profiles_photo_url"`
	BirthData    *time.Time `gorm:"type:date"`
	Occupation   *string    `gorm:"type:varchar(64)"`
	Organization *string    `gorm:"type:varchar(64)"`
	Gender       *Gender    `gorm:"type:char(1)"`
	CreatedAt    time.Time  `gorm:"type:timestamptz;not null;autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"type:timestamptz;not null;autoUpdateTime"`
	// Relation
	UserID uuid.UUID `gorm:"type:uuid;uniqueIndex;not null"`
}
