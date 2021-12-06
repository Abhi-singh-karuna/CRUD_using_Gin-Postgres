package main

import (
	"log"

	"github.com/abhishek-singh/database"
	"github.com/abhishek-singh/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {

	app := gin.Default()

	//load the .env file
	godotenv.Load()

	//connect with databse
	database.ConnectDB()

	router.Routes(app)

	log.Fatal(app.Run(":9000"))
}
