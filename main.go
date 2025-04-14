package main

import (
	"fmt"
	"todolist/config"
	"todolist/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Menginisialisasi koneksi database
	config.ConnectDatabase()

	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r)

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}