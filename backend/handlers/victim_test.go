package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MartinezPosnerValery/Proyecto3/db"
	"github.com/MartinezPosnerValery/Proyecto3/handlers"
	"github.com/MartinezPosnerValery/Proyecto3/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to in-memory database")
	}
	db.DB = database
	if err := db.DB.AutoMigrate(&models.Victim{}); err != nil {
		panic("failed to migrate models")
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/victims", handlers.RegisterVictim)
	router.GET("/victims", handlers.GetAllVictims)
	router.PUT("/victims/:id/details", handlers.UpdateVictimDetails)
	return router
}

func TestVictimEndpoints(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupTestDB()
	router := setupRouter()

	t.Run("registro exitoso con causa", func(t *testing.T) {
		victim := models.Victim{
			ImageURL:     "http://example.com/image.jpg",
			CauseOfDeath: "Sobredosis de memes",
			Details:      "Detalles opcionales",
			FullName:     "Juan Test",
		}

		body, _ := json.Marshal(victim)
		req, _ := http.NewRequest(http.MethodPost, "/victims", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Víctima registrada exitosamente")
		assert.Contains(t, rec.Body.String(), "Sobredosis de memes")
	})

	t.Run("registro sin causa de muerte usa defecto", func(t *testing.T) {
		victim := models.Victim{
			FullName: "Carlos Corazón",
			ImageURL: "http://imagen.com/foto.jpg",
		}

		body, _ := json.Marshal(victim)
		req, _ := http.NewRequest(http.MethodPost, "/victims", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		assert.Contains(t, rec.Body.String(), "Ataque al corazón")
	})

	t.Run("registro sin imagen falla", func(t *testing.T) {
		victim := models.Victim{
			FullName:     "Sin Imagen",
			CauseOfDeath: "Explosión de ideas",
		}

		body, _ := json.Marshal(victim)
		req, _ := http.NewRequest(http.MethodPost, "/victims", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "Debe subir una imagen para que la muerte funcione")
	})

	t.Run("registro con URL inválida falla", func(t *testing.T) {
		victim := models.Victim{
			FullName:     "URL Rota",
			ImageURL:     "not-a-url",
			CauseOfDeath: "Ruido blanco",
		}

		body, _ := json.Marshal(victim)
		req, _ := http.NewRequest(http.MethodPost, "/victims", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "La URL de la imagen no es válida")
	})

	t.Run("actualizar detalles luego de causa", func(t *testing.T) {
		victim := models.Victim{
			FullName:     "Detalle Guy",
			CauseOfDeath: "Ahogo",
			ImageURL:     "http://example.com/img.jpg",
		}
		db.DB.Create(&victim)

		update := map[string]string{"details": "En la piscina municipal"}
		body, _ := json.Marshal(update)
		url := fmt.Sprintf("/victims/%d/details", victim.ID)

		req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Detalles actualizados exitosamente")
	})

	t.Run("listar víctimas", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/victims", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var victims []models.Victim
		err := json.Unmarshal(rec.Body.Bytes(), &victims)
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(victims), 1)
	})
}
