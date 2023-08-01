package sqlite

import (
	"database/sql"

	"github.com/josiahdenton/startme/internal/domain/starter"
)

type SqliteDatabase struct {
	conn *sql.DB
}

func (sDb *SqliteDatabase) Connect(path string) error {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	sDb.conn = conn
	return nil
}

func (sDb *SqliteDatabase) Setup() {
    q := `

    `
}

func (sDb *SqliteDatabase) Close() {
    sDb.conn.Close()
}

func (sDb *SqliteDatabase) Save(starter starter.Starter) error {
	return nil
}

func (sDb *SqliteDatabase) All() []starter.Starter {
	return nil
}
