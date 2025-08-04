package model

type Role struct {
	// ID role sekaligus nama
	ID string `gorm:"type:varchar(32);primaryKey"` // ID role sekaligus nama role
	// Deskripsi role
	Description string `gorm:"type:varchar(255);not null"`
	// User dengan role ini
	Users []*User `gorm:"many2many:user_roles"`
	// Izin role ini
	Permissions []*Permission `gorm:"many2many:role_permissions"`
}
