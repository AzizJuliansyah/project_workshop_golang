package main

import (
	"log"
	"project_workshop_golang_test/config"
	"project_workshop_golang_test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.InitViper()
	if err != nil {
		log.Fatal("Error while configure viper, " + err.Error())
	} else  {
		log.Println("Success configure viper")
	}


	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error while configure to mysql", err.Error())
	} else {
		log.Println("Successfully connect to db")
	}
	

	
	router := gin.Default()
	routes.Routes(router, db)


	router.Run(":8000") // default port :8080
}