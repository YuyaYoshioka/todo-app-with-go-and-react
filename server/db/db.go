package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/YuyaYoshioka/todo-app-with-go-and-react/server/entity"

)

var (
	db *gorm.DB
	err error
)

const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "root"

	// パスワード
	DBPass = ""

	// プロトコル
	DBProtocol = "tcp(127.0.0.1:3306)"

	// DB名
	DBName = "go_react_todo"
)

func AutoMigration() {
	db.AutoMigrate(&entity.Todo{})
}

func DBOpen() *gorm.DB{
	connectTemplate := "%s:%s@%s/%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)
	db, err := gorm.Open(Dialect, connect + "?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}

func Init() {
	db = DBOpen()
	AutoMigration()
}

func GetDB() *gorm.DB{
    return db
}

func Close() {
    if err := db.Close(); err != nil {
        panic(err)
    }
}
