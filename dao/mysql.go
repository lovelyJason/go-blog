package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql" // 注意，这里必须要有
)

type MyDB struct {
	*sql.DB
}

var DB *sql.DB

// 注意：这里是一个简单的连接池，如果需要扩展更多的功能，需要自定义结构体
var DB1 *MyDB

func (m *MyDB) CustomQuery(query string) (*sql.Rows, error) {
	// 你可以在 Mdb 上添加更多的方法
	return m.Query(query)
}

func init() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	loc := url.QueryEscape("Asia/Shanghai")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true",
		user, password, host, port, name, loc)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Printf("链接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	//最大连接数，默认不配置，是不限制最大连接数
	db.SetMaxOpenConns(100)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 3)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)

	}
	DB = db

}
