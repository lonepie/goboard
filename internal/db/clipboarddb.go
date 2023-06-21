package db

import (
	"database/sql"

	m "github.com/lonepie/goboard/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

// ClipboardEntry represents an entry in the clipboard.
// type ClipboardEntry struct {
// 	ID        string    `json:"ID"`
// 	Data      string    `json:"Data"`
// 	Timestamp time.Time `json:"Timestamp"`
// 	RowID     int       `json:"RowID"`
// }

// ClipboardDB represents the clipboard database.
type ClipboardDB struct {
	*sql.DB
}

var clipboardDB *ClipboardDB

// InitClipboardDB creates a new ClipboardDB instance.
func InitClipboardDB(dbPath string) (*ClipboardDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create the clipboard table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clipboard (
		id TEXT NOT NULL PRIMARY KEY,
		data TEXT NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		return nil, err
	}

	clipboardDB = &ClipboardDB{db}
	defer clipboardDB.Close()

	return clipboardDB, nil
}

func GetClipboardDB() *ClipboardDB {
	// if clipboardDB != nil {
	// 	return clipboardDB
	// }
	// cdb, _ := NewClipboardDB("clipboard.db")
	// return cdb
	// return nil
	return clipboardDB
}

// Save inserts or replaces a ClipboardEntry in the ClipboardDB
func (cdb *ClipboardDB) Save(entry *m.ClipboardEntry) error {
	_, err := cdb.Exec("INSERT OR REPLACE INTO clipboard (id, data, timestamp) VALUES (?, ?, ?)", entry.ID, entry.Data, entry.Timestamp)
	return err
}

// ReadEntries reads all clipboard entries.
func (cdb *ClipboardDB) ReadEntries() ([]*m.ClipboardEntry, error) {
	rows, err := cdb.Query("SELECT id, data, timestamp, rowid FROM clipboard ORDER BY rowid DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := []*m.ClipboardEntry{}
	for rows.Next() {
		entry := &m.ClipboardEntry{}
		err := rows.Scan(&entry.ID, &entry.Data, &entry.Timestamp, &entry.RowID)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// GetEntry gets a specific ClipboardEntry by RowID
func (cdb *ClipboardDB) GetEntry(rowID int) (*m.ClipboardEntry, error) {
	rows, err := cdb.Query("SELECT id, data, timestamp, rowid FROM clipboard WHERE rowid = ?", rowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entry := &m.ClipboardEntry{}
	for rows.Next() {
		rows.Scan(&entry.ID, &entry.Data, &entry.Timestamp, &entry.RowID)
	}
	return entry, nil
}

// UpdateEntry updates an existing clipboard entry.
func (cdb *ClipboardDB) UpdateEntry(entry m.ClipboardEntry) error {
	_, err := cdb.Exec("UPDATE clipboard SET data = ?, timestamp = ? WHERE id = ?", entry.Data, entry.Timestamp, entry.ID)
	return err
}

// DeleteEntry deletes a clipboard entry by ID.
func (cdb *ClipboardDB) DeleteEntry(id string) error {
	_, err := cdb.Exec("DELETE FROM clipboard WHERE id = ?", id)
	return err
}
