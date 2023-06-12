package clipboardmonitor

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// ClipboardEntry represents an entry in the clipboard.
type ClipboardEntry struct {
	ID        string
	Data      string
	Timestamp time.Time
	RowID     int
}

// ClipboardDB represents the clipboard database.
type ClipboardDB struct {
	*sql.DB
}

// NewClipboardDB creates a new ClipboardDB instance.
func NewClipboardDB(dbPath string) (*ClipboardDB, error) {
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

	return &ClipboardDB{db}, nil
}

// Close closes the clipboard database connection.
// func (cdb *ClipboardDB) Close() error {
// 	return cdb.Close()
// }

// CreateEntry creates a new clipboard entry.
// func (cdb *ClipboardDB) CreateEntry(data string) error {
// 	_, err := cdb.db.Exec("INSERT INTO clipboard (data) VALUES (?)", data)
// 	return err
// }

// Save inserts or replaces a ClipboardEntry in the ClipboardDB
func (cdb *ClipboardDB) Save(entry *ClipboardEntry) error {
	_, err := cdb.Exec("INSERT OR REPLACE INTO clipboard (id, data, timestamp) VALUES (?, ?, ?)", entry.ID, entry.Data, entry.Timestamp)
	return err
}

// ReadEntries reads all clipboard entries.
func (cdb *ClipboardDB) ReadEntries() ([]*ClipboardEntry, error) {
	rows, err := cdb.Query("SELECT id, data, timestamp, rowid FROM clipboard")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := []*ClipboardEntry{}
	for rows.Next() {
		entry := &ClipboardEntry{}
		err := rows.Scan(&entry.ID, &entry.Data, &entry.Timestamp, &entry.RowID)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// GetEntry gets a specific ClipboardEntry by RowID
func (cdb *ClipboardDB) GetEntry(rowID int) (*ClipboardEntry, error) {
	rows, err := cdb.Query("SELECT id, data, timestamp, rowid FROM clipboard WHERE rowid = ?", rowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	entry := &ClipboardEntry{}
	for rows.Next() {
		rows.Scan(&entry.ID, &entry.Data, &entry.Timestamp, &entry.RowID)
	}
	return entry, nil
}

// UpdateEntry updates an existing clipboard entry.
// func (cdb *ClipboardDB) UpdateEntry(entry ClipboardEntry) error {
// 	_, err := cdb.db.Exec("UPDATE clipboard SET data = ?, timestamp = ? WHERE id = ?", entry.Data, entry.Timestamp, entry.ID)
// 	return err
// }

// // DeleteEntry deletes a clipboard entry by ID.
// func (cdb *ClipboardDB) DeleteEntry(id int) error {
// 	_, err := cdb.db.Exec("DELETE FROM clipboard WHERE id = ?", id)
// 	return err
// }
