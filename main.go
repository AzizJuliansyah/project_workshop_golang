package main

import (
	"fmt"
	"log"
	"project_workshop_golang_test/config"
	"project_workshop_golang_test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitViper()
	if err != nil {
		log.Fatal("Gagal memuat konfigurasi viper, " + err.Error())
	} else {
		fmt.Println("Berhasil memuat konfigurasi viper")
	}

	db, err := config.InitDatabase()
	if err != nil {
		log.Fatal("Gagal memuat konfigurasi database, " + err.Error())
	} else {
		fmt.Println("Berhasil memuat konfigurasi database")
	}

	router := gin.Default()
	routes.Routes(router, db)
	router.Run(":9000")
}