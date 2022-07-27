package database

import (
	"grpcCRUD/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB

/*
func connect(connectionString string) error {
	var err error

	connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return err
	}
	log.Printf("Connection Successful...")
	connector.AutoMigrate(&model.Book{})
	return nil
}

func init() {
	330config := Config{
		ServerName: "localhost:6",
		User:       "root",
		Password:   "TFT@us@123",
		DB:         "book",
	}

	connectionString := getConnectionString(config)
	err := connect(connectionString)
	if err != nil {
		log.Panic(err.Error())
	}
}*/

var dbURL = "root:TFT@us@123@tcp(127.0.0.1:3306)/book"

var C *gorm.DB

func init() {

	connector, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	//connector.Exec("CREATE DATABASE IF NOT EXISTS" + "book")

	log.Printf("Connection Successful...")
	connector.AutoMigrate(&model.Book{})
	C = connector

	// connector, err := gorm.Open(sqlite.Open("C:/"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Faild to Connect to the database")
	// }

	// connector.AutoMigrate(&model.Book{})
	// connector.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Book{})

}
