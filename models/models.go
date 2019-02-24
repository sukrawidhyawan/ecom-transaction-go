package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var dbUsername string
var dbPassword string
var dbName string
var dbHost string

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	dbUsername = os.Getenv("db_user")
	dbPassword = os.Getenv("db_pass")
	dbName = os.Getenv("db_name")
	dbHost = os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUsername, dbName, dbPassword) //Build connection string
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println(err)
	}
	conn.AutoMigrate(&Cart{}, &Order{})
	db = conn
}

func GetDb() *gorm.DB {
	return db
}
