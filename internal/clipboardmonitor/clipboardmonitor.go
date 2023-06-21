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

// InitClipboardMonitor creates a new ClipboardDB instance and starts monitoring the clipboard.
func InitClipboardMonitor(dbPath string) (*ClipboardMonitor, error) {
	db, err := db.InitClipboardDB(dbPath)
	if err != nil {
		return nil, err
	}

	entryChan := make(chan m.ClipboardEntry)
	monitor := &ClipboardMonitor{DB: db, EntryChan: entryChan}

	err = clipboard.Init()
	if err != nil {
		log.Fatalf("Error initializing clipboard: %v", err)
		// log.Println("Error initializing clipboard:", err)
	}

	// Start monitoring the clipboard
	go monitor.monitorClipboard()

	return monitor, nil
}

func hash(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
	// h := md5.New()
	// h.Write([]byte(data))
	// return hex.EncodeToString(h.Sum(nil))
}

// monitorClipboard monitors the clipboard for changes and inserts new entries into the database.
func (monitor *ClipboardMonitor) monitorClipboard() {
	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range ch {
		strData := string(data)
		entry := m.ClipboardEntry{
			ID:        hash(strData),
			Data:      strData,
			Timestamp: time.Now(),
		}
		monitor.DB.Save(&entry)
		monitor.EntryChan <- entry
	}
}

// func (entry *ClipboardEntry) WriteToClipboard() {
// 	clipboard.Init()
// 	clipboard.Write(clipboard.FmtText, []byte(entry.Data))
// 	log.Println("Wrote entry to clipboard:", entry.RowID)
// }
