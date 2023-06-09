package main

import (
	"connorlucier/match-picker/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=mpadmin password=mpadmin dbname=matchpicker port=25432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec(`create extension if not exists "uuid-ossp";`)
	tryMigrate(db, model.Team{})
	tryMigrate(db, model.League{})
	tryMigrate(db, model.Season{})
	tryMigrate(db, model.Match{})
	tryMigrate(db, model.Pick{})
	tryMigrate(db, model.User{})
	tryMigrate(db, model.Group{})
}

func tryMigrate(db *gorm.DB, t interface{}) {
	err := db.AutoMigrate(&t)
	if err != nil {
		panic(err)
	}
}
