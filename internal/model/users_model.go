package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Email        string         `gorm:"type:varchar(255);not null;index:idx_users_email;uniqueIndex:uidx_users_email_not_deleted,where:deleted_at IS NULL"`
	PasswordHash string         `gorm:"type:varchar(255);not null"`
	Role         []*Role        `gorm:"many2many:user_roles"`
	CreatedAt    time.Time      `gorm:"type:timestamptz;not null;autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"type:timestamptz;not null;autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"type:timestamptz;index"`
}
