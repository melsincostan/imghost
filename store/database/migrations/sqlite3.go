//go:build !nosqlite

package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed sqlite3/*
var _sqlite3_migrations embed.FS

type Sqlite struct {
	db *sql.DB
}

func (s *Sqlite) Up() (applied int, err error) {
	source := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: _sqlite3_migrations,
	}

	applied, err = migrate.Exec(s.db, "sqlite3", source, migrate.Up)
	if err != nil {
		return applied, fmt.Errorf("sqlite3 migration: %w", err)
	}
	return
}

func NewSqlite(db *sql.DB) *Sqlite {
	return &Sqlite{
		db: db,
	}
}
