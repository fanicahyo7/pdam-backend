package migrations

import (
	"fmt"
	"pdam/model"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("ini migration")
}

func UserMigration(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.User{}) {
		db.Migrator().CreateTable(&model.User{})
	}

}

func UserRollback(db *gorm.DB) {
	if db.Migrator().HasTable(&model.User{}) {
		db.Migrator().DropTable(&model.User{})
	}
}
