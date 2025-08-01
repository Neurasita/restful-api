package model

import "gorm.io/gorm"

func AutoMigrate(tx *gorm.DB) error {
	if err := tx.AutoMigrate(
		&User{},
		&Role{},
		&Permission{},
		&Profiles{},
	); err != nil {
		return err
	}
	return nil
}
