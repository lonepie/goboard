package clipboardmonitor

import (
	"context"
	"crypto/md5"
	"io"
	"log"
	"time"

	"golang.design/x/clipboard"
)

type ClipboardMonitor struct {
	DB        *ClipboardDB
	EntryChan chan ClipboardEntry
}

// NewClipboardMonitor creates a new ClipboardDB instance and starts monitoring the clipboard.
func NewClipboardMonitor(dbPath string) (*ClipboardMonitor, error) {
	db, err := NewClipboardDB(dbPath)
	if err != nil {
		return nil, err
	}

	entryChan := make(chan ClipboardEntry)
	monitor := &ClipboardMonitor{DB: db, EntryChan: entryChan}

	// Start monitoring the clipboard
	go monitor.monitorClipboard()

	return monitor, nil
}

func hash(data string) []byte {
	h := md5.New()
	io.WriteString(h, data)
	return h.Sum(nil)
}

// monitorClipboard monitors the clipboard for changes and inserts new entries into the database.
func (monitor *ClipboardMonitor) monitorClipboard() {
	err := clipboard.Init()
	if err != nil {
		log.Println("Error initializing clipboard:", err)
	}

	ch := clipboard.Watch(context.Background(), clipboard.FmtText)
	for data := range ch {
		//fmt.Println(data)
		strData := string(data)
		entry := ClipboardEntry{
			ID:        hash(strData),
			Data:      strData,
			Timestamp: time.Now(),
		}
		monitor.DB.Save(&entry)
		monitor.EntryChan <- entry
	}
}
