package storage

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SqliteNewConnectionWithFileName(fileName *string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("data/%s.db", *fileName)
	return newSqliteConnection(&dsn)
}

func SqliteNewConnectionWithInMemory() (*gorm.DB, error) {
	dsn := "file::memory:?cache=shared"
	return newSqliteConnection(&dsn)
}

func newSqliteConnection(dsn *string) (*gorm.DB, error) {

	dial := sqlite.Open(*dsn)

	db, conErr := newConnection(dial)
	if conErr != nil {
		return db, conErr
	}

	return db, nil
}
