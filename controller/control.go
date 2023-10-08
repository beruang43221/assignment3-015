package controller

import (
	"github.com/beruang43221/assignment3-015/database"
	"github.com/beruang43221/assignment3-015/models"
	"github.com/gin-gonic/gin"
)

func UpdateWeather(c *gin.Context){
	db := database.GetDB()
	var micro models.Microservice

	id := c.Param("id")

	data := db.First(&micro, id)

	if data.RowsAffected == 0 {
		//Jika data dengan ID = 1 tidak ditemukan
		createData := models.Microservice{
			ID: 1, 
			Water: 0, 
			Wind: 0}
		db.Create(&createData)
		// log.Printf("{\n    \"water\": %d,\n    \"wind\": %d\n}\nWater Status: %s\nWind Status: %s",
    //         micro.Water, micro.Wind, getWaterStatus(micro.Water), getWindStatus(micro.Wind))
	}else {
		// Jika data dengan ID = 1 ditemukan
		var newData models.Microservice
		c.BindJSON(&newData)
		micro.Water = newData.Water
		micro.Wind = newData.Wind
		db.Save(&micro)
		// log.Printf("{\n    \"water\": %d,\n    \"wind\": %d\n}\nWater Status: %s\nWind Status: %s",
    //         micro.Water, micro.Wind, getWaterStatus(micro.Water), getWindStatus(micro.Wind))
	}
}


