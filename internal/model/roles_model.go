package model

type Role struct {
	ID          string        `gorm:"type:varchar(32);primaryKey"` // ID role sekaligus nama role
	Description string        `gorm:"type:varchar(255);not null"`
	Users       []*User       `gorm:"many2many:user_roles"`
	Permissions []*Permission `gorm:"many2many:role_permissions"`
}
