package db
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// Init 初始化数据库
func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/store")
	if err != nil {
		fmt.Println(err)
		return
	}
	db = database
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")
}

// 指针指向db
func GetDB() *sqlx.DB {
    return db
}