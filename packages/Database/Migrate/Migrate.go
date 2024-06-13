package Migrate

import (
	Location "forro_project/api/Location/Domain/Model"
	"forro_project/packages/Database/Config"
	"gorm.io/gorm"
)

func RunMigrate() {
	db := Config.StartDbConnection()
	locationMigrate(db)
}

func locationMigrate(db *gorm.DB) {
	db.AutoMigrate(&Location.Country{}, &Location.State{}, &Location.City{}, &Location.Address{})
}
