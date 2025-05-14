package main

import (
	"log"

	"github.com/MartinezPosnerValery/Proyecto3/db"
	"github.com/MartinezPosnerValery/Proyecto3/handlers"
	"github.com/MartinezPosnerValery/Proyecto3/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = models.Victim{}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	db.InitDB()
	db.DB.AutoMigrate(&models.Victim{})

	r := gin.Default()

	// Middleware para permitir CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/victims", handlers.RegisterVictim)
	r.GET("/victims", handlers.GetAllVictims)

	println("Servidor backend corriendo en http://0.0.0.0:8080")
	r.Run("0.0.0.0:8080")
}
