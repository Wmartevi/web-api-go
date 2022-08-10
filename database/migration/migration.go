package migration

import (
	"github.com/Wmartevi/web-api-go/model"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(model.Books{})
}
