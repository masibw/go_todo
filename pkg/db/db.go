package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)
//DB接続
func GormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	DBHOST := os.Getenv("DBHOST")
	DBNAME := os.Getenv("DBNAME")

	CONNECT := DBUSER + ":" + DBPASS + "@("+DBHOST+")/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open("mysql", CONNECT)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}