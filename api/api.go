package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lonepie/goboard/frontend"
	"github.com/lonepie/goboard/internal/db"
	m "github.com/lonepie/goboard/internal/model"
)

type ClipboardAPI struct {
	DB     *db.ClipboardDB
	Router *gin.Engine
}

var clipboardAPI *ClipboardAPI

func StartAPI(dbPath string, staticFiles string) {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}

	//initialize the database
	cbdb, err := db.InitClipboardDB(dbPath)
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}

	clipboardAPI = &ClipboardAPI{DB: cbdb, Router: router}

	router.Use(cors.New(corsConfig))

	if len(staticFiles) > 0 {
		log.Printf("Serving static files from path: %v\n", staticFiles)
		router.Use(static.Serve("/", static.LocalFile(staticFiles, false)))
	} else {
		log.Println("Serving static files from embedded fs")
		router.Use(static.Serve("/", frontend.GetFS(true)))
	}
	// router.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))

	api := router.Group("/api")
	{
		api.GET("/clipboard", getAllClipboardEntries)
		api.PUT("clipboard/:id", updateClipboardEntry)
		api.DELETE("/clipboard/:id", deleteClipboardEntry)
	}

	router.Run(":3000")
}

func getAllClipboardEntries(c *gin.Context) {
	// db := clipboardmonitor.GetClipboardDB()
	entries, err := clipboardAPI.DB.ReadEntries()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entries)
}

func updateClipboardEntry(c *gin.Context) {
	id := c.Param("id")
	var entry m.ClipboardEntry
	if err := c.ShouldBindJSON(&entry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entry.ID = id
	entry.Timestamp = time.Now()

	// err := clipboardmonitor.GetClipboardDB().UpdateEntry(entry)
	err := clipboardAPI.DB.UpdateEntry(entry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, entry)
}

func deleteClipboardEntry(c *gin.Context) {
	id := c.Param("id")
	// err := clipboardmonitor.GetClipboardDB().DeleteEntry(id)
	err := clipboardAPI.DB.DeleteEntry(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Clipboard entry %v deleted successfully", id)})
}
