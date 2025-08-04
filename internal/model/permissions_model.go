package model

type Permission struct {
	// Permission id sekaligus nama permission
	// ikuti konvensi app ini dimana permission berisi
	// nama tabel dan aksi. contohnya `users:write`
	ID string `gorm:"type:varchar(32);primaryKey"`
	// Pastikan selalu menambahkan deskripsi izin untuk permission
	Description string `gorm:"type:varchar(255);not null"`
	// Relasi many to many ke tabel Role
	// ! pastikan nama table sama role_permissions pada model Role
	Roles []*Role `gorm:"many2many:role_permissions"`
}
