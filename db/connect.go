package db

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Singleton struct {
	Db *gorm.DB
}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
		dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC`,
			viper.GetString("db.host"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.name"),
			viper.GetInt("db.Port"),
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance.Db = db
	})
	return instance
}
