package api

import (
	"log"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lonepie/goboard/internal/clipboardmonitor"
)

func StartAPI() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	api := router.Group("/api")
	{
		api.GET("clipboard", getAllClipboardEntries)

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
