package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lonepie/goboard/internal/clipboardmonitor"
)

func StartAPI(dbPath string) {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}

	//initialize the database
	clipboardmonitor.NewClipboardDB(dbPath)

	router.Use(cors.New(corsConfig))
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	api := router.Group("/api")
	{
		api.GET("/clipboard", getAllClipboardEntries)
		api.PUT("clipboard/:id", updateClipboardEntry)
		api.DELETE("/clipboard/:id", deleteClipboardEntry)
	}

	router.Run(":3000")
}

func getAllClipboardEntries(c *gin.Context) {
	// db, err := clipboardmonitor.NewClipboardDB("clipboard.db")
	// if err != nil {
	// 	log.Println("Error: ", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }
	db := clipboardmonitor.GetClipboardDB()
	entries, _ := db.ReadEntries()
	c.JSON(http.StatusOK, entries)
}

func updateClipboardEntry(c *gin.Context) {
	id := c.Param("id")
	var entry clipboardmonitor.ClipboardEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry.ID = id
	entry.Timestamp = time.Now()

	err := clipboardmonitor.GetClipboardDB().UpdateEntry(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func deleteClipboardEntry(c *gin.Context) {
	id := c.Param("id")
	err := clipboardmonitor.GetClipboardDB().DeleteEntry(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Clipboard entry %v deleted successfully", id)})
}
