package main

import (
	"github.com/beruang43221/assignment3-015/database"
	"github.com/beruang43221/assignment3-015/routers"
	"github.com/beruang43221/assignment3-015/service"
)

func main() {
	database.StartDB()
	
	go service.HitAPI()
	
	router := routers.SetupRouter()
	
	router.Run(":8083")

}