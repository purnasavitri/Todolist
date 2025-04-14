package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //menyimpan koneksi database 

func ConnectDatabase() {
	dsn := "root:@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local" //konfigurasi koneksi database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal koneksi ke database")
		panic(err)
	}
	fmt.Println("Berhasil koneksi ke database")
}