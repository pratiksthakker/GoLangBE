package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// OpenDbConnection Opening a database and save the reference to `Database` struct.
func OpenDbConnection() *gorm.DB {

	// dialect := "mysql"
	// username := "root"
	// password := "root"
	// dbName := "pacedb"
	// host := "localhost:3306"
	var db *gorm.DB
	var err error

	//databaseURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ", host, username, password, dbName)
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/paceecommercedb")

	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

// GetDb function
func GetDb() *gorm.DB {
	return DB
}
