package main

import (
	"final-task/router"
	"final-task/database"
	"log"
)

func main() {
	// Inisialisasi koneksi database
	database.InitDB()

	// Setup router
	r := router.setupRouter()

	// Menjalankan server
	err := r.Run()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

