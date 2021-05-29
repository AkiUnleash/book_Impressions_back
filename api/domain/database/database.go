package database

import (
	"fmt"
	"os"
	"srb/domain/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Connet() {

	// ローカルで実行時に[Config.ini]を使用する場合。
	var DBMS string
	var USER string
	var PASS string
	var PROTOCOL string
	var DBNAME string

	// if os.Getenv("ENVIROMENT") == "production" {
	// Heroku用
	DBMS = os.Getenv("DB_HOST")
	USER = os.Getenv("DB_USERNAME")
	PASS = os.Getenv("DB_PASSWORD")
	PROTOCOL = os.Getenv("DB_PROTOCOL")
	DBNAME = os.Getenv("DB_NAME")
	// } else {
	// 	// ローカルで実行時に[Config.ini]を使用する場合。
	// 	DBMS = config.Config.Dbms
	// 	USER = config.Config.User
	// 	PASS = config.Config.Pass
	// 	PROTOCOL = config.Config.Protocol
	// 	DBNAME = config.Config.Dbname
	// }

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"

	connection, err := gorm.Open(DBMS, CONNECT)

	fmt.Println(CONNECT)

	if err != nil {
		panic("could not connet to the database")
	}

	DB = connection

	// テーブル自動生成
	connection.AutoMigrate(&models.Account{})
	connection.AutoMigrate(&models.Impression{})
}
