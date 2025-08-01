package model

type Permission struct {
	// Permission id sekaligus nama permission
	// ikuti konvensi app ini dimana permission berisi
	// nama tabel dan aksi. contohnya `users:write`
	ID          string  `gorm:"type:varchar(32);primaryKey"`
	Description string  `gorm:"type:varchar(255);not null"`
	Roles       []*Role `gorm:"many2many:role_permissions"`
}
