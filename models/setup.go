package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() {
	db, err := gorm.Open(mysql.Open("root:123456@/go_coba"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Mahasiswa{})
	DB = db
}
