package models

import "muxapi/database"

func Migrations() {
	db := database.Database
	db.AutoMigrate(&UserCategory{}, &Product{}, &Picture{}, &User{}, &Profile{})
}
