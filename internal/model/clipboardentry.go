package model

import (
	"time"
)

// ClipboardEntry represents an entry in the clipboard.
type ClipboardEntry struct {
	ID        string    `json:"ID"`
	Data      string    `json:"Data"`
	Timestamp time.Time `json:"Timestamp"`
	RowID     int       `json:"RowID"`
}
