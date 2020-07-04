package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)


// DBの初期化
func dbInit() {
	db := gormConnect()
	// コネクション解放
	defer db.Close()
}

func gormConnect() *gorm.DB {
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