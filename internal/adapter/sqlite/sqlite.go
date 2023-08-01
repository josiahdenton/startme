package sqlite

import (
	"database/sql"
	"log"

	"github.com/josiahdenton/startme/internal/domain/starter"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDatabase struct {
	conn *sql.DB
}

func New(path string) SqliteDatabase {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Error: failed to setup table: %v", err)
	}
	return SqliteDatabase{
		conn: conn,
	}
}

// Setup will create the DB table if it does not yet exist.
//
// This should be called before making any R/W on the db.
func (sDb *SqliteDatabase) Setup() {
	q := `
       CREATE TABLE IF NOT EXISTS starters (
           id INTEGER PRIMARY KEY,
           title TEXT,
           content TEXT
       ) 
    `
	_, err := sDb.conn.Exec(q)
	if err != nil {
		log.Fatalf("Error: failed to setup table: %v", err)
	}
}

func (sDb *SqliteDatabase) Close() {
	sDb.conn.Close()
}

func (sDb *SqliteDatabase) Save(title, content string) error {
	q := `INSERT INTO starters (id,title,content) VALUES (NULL,?,?)`
	_, err := sDb.conn.Exec(q, title, content)
	if err != nil {
		log.Fatalf("Error: failed to insert starter")
	}
	return nil
}

func (sDb *SqliteDatabase) All() []starter.Starter {
	q := `SELECT title, content FROM starters`
	rows, err := sDb.conn.Query(q)
	if err != nil {
		log.Fatalf("Error: failed to query DB: %v", err)
	}
	defer rows.Close()

	starters := make([]starter.Starter, 0, 10)
	for rows.Next() {
		var (
			title   string
			content string
		)
		rows.Scan(&title, &content)
		starters = append(starters, starter.Starter{
			Title:   title,
			Content: content,
		})
	}
	return starters
}
