package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MySQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// type Database struct {
// 	db *gorm.DB
// }

// NewMySQLDBConnection return instance of DB Connection
func NewMySQLDBConnection(c *MySQLConfig) (db *gorm.DB, err error) {
	// network name default tcp optional as long as host and port wrapped in ()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?timeout=3s&charset=utf8mb4&parseTime=true&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Connect Database failed: %s", err.Error())
		return
	}

	db.NamingStrategy = schema.NamingStrategy{SingularTable: true}

	// database = &Database{db}

	log.Println("MySQL Database connected")

	return
}
