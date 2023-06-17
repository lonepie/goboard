package api

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lonepie/goboard/internal/clipboardmonitor"
)

func StartAPI() {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	api := router.Group("/api")
	{
		api.GET("/clipboard", getAllClipboardEntries)
		// api.PUT("clipboard/:id", updateClipboardEntry)
	}

	// router.Static("/frontend", "./frontend/dist")

	router.Run(":3000")
}

func getAllClipboardEntries(c *gin.Context) {
	db, err := clipboardmonitor.NewClipboardDB("clipboard.db")
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	entries, _ := db.ReadEntries()
	c.JSON(http.StatusOK, entries)
}
