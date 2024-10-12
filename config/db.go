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

func NewDatabase(db Database) *Database {
	return &Database{
		Host: db.Host,
		User: db.User,
		Pass: db.Pass,
		Name: db.Name,
		Port: db.Port,
	}
}

func (d *Database) Connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		d.Host,
		d.User,
		d.Pass,
		d.Name,
		d.Port,
	)

	return gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
}
