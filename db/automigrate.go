package db

import "github.com/srazap/luckbuy/models"

func MigrateModels() {
	if err := GetInstance().Db.AutoMigrate(
		&models.User{},
		&models.Session{},
	); err != nil {
		panic(err)
	}
}
