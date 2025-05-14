package handlers

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/MartinezPosnerValery/Proyecto3/db"
	"github.com/MartinezPosnerValery/Proyecto3/models"
	"github.com/gin-gonic/gin"
)

func RegisterVictim(c *gin.Context) {
	var victim models.Victim
	if err := c.ShouldBindJSON(&victim); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar URL de imagen
	_, err := url.ParseRequestURI(victim.ImageURL)
	if victim.ImageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Debe subir una imagen para que la muerte funcione"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La URL de la imagen no es válida"})
		return
	}

	// Si no se proporciona una causa de muerte, usar valor por defecto
	if victim.CauseOfDeath == "" {
		victim.CauseOfDeath = "Ataque al corazón"
	}

	// Establecer la marca de tiempo de la muerte
	victim.DeathTimestamp = time.Now().Unix() + 40

	victim.IsDead = true

	// Crear la víctima en la base de datos
	if err := db.DB.Create(&victim).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar a la víctima", "details": err.Error()})
		return
	}

	// Devolver la respuesta que incluya la causa de muerte
	c.JSON(http.StatusCreated, gin.H{"message": "Víctima registrada exitosamente", "causeOfDeath": victim.CauseOfDeath})
}

func GetAllVictims(c *gin.Context) {
	var victims []models.Victim
	db.DB.Find(&victims)
	c.JSON(http.StatusOK, victims)
}

func UpdateVictimDetails(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input struct {
		Details string `json:"details"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	var victim models.Victim
	if err := db.DB.First(&victim, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Víctima no encontrada"})
		return
	}

	victim.Details = input.Details
	db.DB.Save(&victim)

	c.JSON(http.StatusOK, gin.H{"message": "Detalles actualizados exitosamente"})
}
