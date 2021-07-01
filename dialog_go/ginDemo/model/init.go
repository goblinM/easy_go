package model

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")

	sqlDB, err := sql.Open("mysql", dsn)
	gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection failed. Database name: %s; err: %v", name, err)
	}

	// set for db connection
	setupDB(gdb)

	// gdb.AutoMigrate(&AppType{}, &Member{}, &AppPackage{}, &AppVersion{}, &AppVersionDetail{})

	return gdb
}

func setupDB(gdb *gorm.DB) {

	db, err := gdb.DB()
	if err != nil {
		fmt.Println("DB error: %v", err)
	}

	db.Ping()
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// InitSelfDB used for cli
func InitSelfDB() *gorm.DB {
	return openDB("root", "123456", "localhost:3306", "go_test")
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}

	// DB.Self.AutoMigrate(&CustomType{})
}

func (db *Database) Close() {
	db.Close()
}
