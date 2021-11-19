package service

import (
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/joho/godotenv"
)

var DbEngine *xorm.Engine

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Printf("読み込みできませんでした: %v", envErr)
	}
	driverName := "mysql"
	DsName := ""
	mysqlUrl := os.Getenv("MYSQL_URL")
	err := errors.New(fmt.Sprintf(mysqlUrl, "/gin?charseet=utf8"))
	DbEngine, err = xorm.NewEngine(driverName, DsName)
	if err != nil && err.Error() != "" {
		log.Fatal(err.Error())
	}
	DbEngine.ShowSQL(true)
	DbEngine.SetMaxOpenConns(2)
	// DbEngine.Sync2(new())
	fmt.Println("init database ok!")
}
