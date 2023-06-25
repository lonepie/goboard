package clipboardmonitor

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	"time"

	"github.com/lonepie/goboard/internal/db"
	m "github.com/lonepie/goboard/internal/model"
	"golang.design/x/clipboard"
)

type ClipboardMonitor struct {
	DB        *db.ClipboardDB
	EntryChan chan m.ClipboardEntry
}

// NewClipboardMonitor creates a new ClipboardMonitor instance.
func NewClipboardMonitor(dbPath string) (*ClipboardMonitor, error) {
	db, err := db.InitClipboardDB(dbPath)
	if err != nil {
		return nil, err
	}

	entryChan := make(chan m.ClipboardEntry)
	monitor := &ClipboardMonitor{DB: db, EntryChan: entryChan}

	err = clipboard.Init()
	if err != nil {
		log.Fatalf("Error initializing clipboard: %v", err)
	}

	// Start monitoring the clipboard
	// go monitor.monitorClipboard()

	return monitor, nil
}

// MonitorClipboard monitors the clipboard for changes and inserts new entries into the database.
func (monitor *ClipboardMonitor) MonitorClipboard() {
	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range ch {
		strData := string(data)
		entry := m.ClipboardEntry{
			ID:        fmt.Sprintf("%x", md5.Sum([]byte(strData))),
			Data:      strData,
			Timestamp: time.Now(),
		}
		monitor.DB.Save(&entry)
		monitor.EntryChan <- entry
	}
}
