package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host string
	User string
	Pass string
	Name string
	Port string
}

func NewDatabase(db Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		db.Host,
		db.User,
		db.Pass,
		db.Name,
		db.Port,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
