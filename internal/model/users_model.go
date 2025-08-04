package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User credentials table
type User struct {

	// PK

	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	// Main data

	// User email
	Email string `gorm:"type:varchar(255);not null;index:idx_users_email;uniqueIndex:uidx_users_email_not_deleted,where:deleted_at IS NULL"`
	// User password hash, gunakan argon2 untuk algoritma
	// hashing pada password
	PasswordHash string `gorm:"type:varchar(255);not null"`
	// Timestamp information
	CreatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime"`
	// Timestamp information
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;autoUpdateTime"`
	// Timestamp information
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz;index"`

	// Relations

	// One to one relation with profile
	Profile *Profile
	// Many to many relation to Role in table user_roles
	// ! pastikan nama tabel relasi sama dengan Users yang ada pada Role
	Role []*Role `gorm:"many2many:user_roles"`
}
