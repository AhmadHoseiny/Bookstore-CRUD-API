package DBConnector

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// db is the variable that holds the connection to the database
	db *gorm.DB
)

func Connect() {

	dsn := "{username}:{password}@tcp(0.0.0.0:3306)/{databaseName}?charset=utf8&parseTime=True&loc=Local"
	curDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		db = curDB
	}
}

func GetDB() *gorm.DB {
	return db
}
