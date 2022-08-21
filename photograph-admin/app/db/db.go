package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// DBの接続情報を取得
func Init() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	dsn := os.Getenv(("MYSQL_USER")) + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv(("MYSQL_HOST")) + ":" + os.Getenv(("MYSQL_PORT")) + ")/" + os.Getenv(("MYSQL_DATABASE")) + "charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

// DBの接続を閉じる
func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err = sqlDB.Close(); err != nil {
		fmt.Println(err)
	}
}
