package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Database struct {
	Self  *gorm.DB
	Common *gorm.DB
	Storage *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
	    username,
	    password,
	    addr,
	    name,
	    true,
	    "Local",
	)

	db, err := gorm.Open("mysql", config)

	if (err != nil) {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(20)  // setting open connection count.
	db.DB().SetMaxIdleConns(20)  // setting idle connection count.
}

func InitSelfDB() *gorm.DB {
	return openDB(
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
}

func InitCommonDB() *gorm.DB {
	return openDB(
		viper.GetString("common_db.username"),
		viper.GetString("common_db.password"),
		viper.GetString("common_db.addr"),
		viper.GetString("common_db.name"),
	)
}

func InitStorageDB() *gorm.DB {
	return openDB(
		viper.GetString("storage_db.username"),
		viper.GetString("storage_db.password"),
		viper.GetString("storage_db.addr"),
		viper.GetString("storage_db.name"),
	)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func GetCommonDB() *gorm.DB {
	return InitCommonDB()
}

func GetStorageDB() *gorm.DB {
	return InitStorageDB()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
		Common: GetCommonDB(),
		Storage: GetStorageDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
	DB.Common.Close()
	DB.Storage.Close()
}