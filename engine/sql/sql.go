package sql

import (
	"fmt"
	"github.com/dytlan/moonlay-todo-list/impl/accessor"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Connection string
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
}

func NewClient(config Config) (db *gorm.DB, err error) {
	var connectionType gorm.Dialector
	switch config.Connection {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Port, config.Name)
		connectionType = mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Host, config.Username, config.Password, config.Name, config.Port)
		connectionType = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("db connection type not found")
	}
	return gorm.Open(connectionType, &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&accessor.ToDoDB{})
}
