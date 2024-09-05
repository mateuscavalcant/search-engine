package controllers

import (
	"log"
	"net/http"
	"search-engine/config/db"
	"search-engine/internal/models"
	"search-engine/pkg/repositories"

	"github.com/gin-gonic/gin"
)

func SearchEngine(c *gin.Context) {
	var req models.SearchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if err := c.ShouldBindJSON(&req); err != nil {
			log.Println("Error JSON:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
			return
		}

		database := db.GetDB()

		searchTerm := req.Search

		users, err := repositories.SearchUsers(database, searchTerm)
		if err != nil {
			log.Println("User not found", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
			return
		}

		resp := map[string]interface{}{
			"users": users,
		}
		c.JSON(http.StatusOK, resp)
	}

}
