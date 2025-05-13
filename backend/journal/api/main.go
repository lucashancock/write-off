package api

import (
	"journal/db"
	"journal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiServer() {
	// correctly handle the error below
	database, err := db.InitDB()
	if err != nil {
		panic("failed to connect to database")
	}

	router := gin.Default()

	// API to create a new journal entry
	router.POST("/journals", func(c *gin.Context) {
		var journal models.Journal
		if err := c.ShouldBindJSON(&journal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.CreateJournal(database, &journal); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create journal"})
			return
		}
		c.JSON(http.StatusCreated, journal)
	})

	// API to get a journal entry by ID
	router.GET("/journals/:id", func(c *gin.Context) {
		var uri struct {
			ID uint `uri:"id" binding:"required"`
		}
		if err := c.ShouldBindUri(&uri); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}

		journal, err := db.GetJournal(database, uri.ID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusOK, journal)
	})

	// API to get all journal entries
	router.GET("/journals", func(c *gin.Context) {
		var journals []models.Journal
		if err := database.Find(&journals).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch journals"})
			return
		}
		c.JSON(http.StatusOK, journals)
	})

	router.Run(":8081")
}
