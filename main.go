package main

import (
	"fmt"
	"net/http"
	"todolist/config"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `gorm:"column:id_user;primaryKey;table:user" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"-"`
}

func (User) TableName() string {
	return "user"
}

func main() {
	// Menginisialisasi koneksi database
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		var user []User
		config.DB.Find(&user)
		c.JSON(http.StatusOK, gin.H{"data": user})
	})

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
