package migrations

import (
	"pdam/model"

	"gorm.io/gorm"
)

func UserMigration(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.User{}) {
		db.Migrator().CreateTable(&model.User{})
	}
}
