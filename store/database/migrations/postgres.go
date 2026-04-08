//go:build !nopostgres

package migrations

import (
	"database/sql"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed postgres/*
var _postgres_migrations embed.FS

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Up() (applied int, err error) {
	source := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: _postgres_migrations,
	}

	applied, err = migrate.Exec(p.db, "postgres", source, migrate.Up)
	if err != nil {
		return applied, fmt.Errorf("postgres migration: %w", err)
	}
	return
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}
