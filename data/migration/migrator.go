package migration

import (
	"database/sql"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const migrationFilesPath = "migrations"

type DBType float64
type MigrationDirection float64

const (
	POSTGRESS DBType = iota
	SQLITE3
)

const (
	UP MigrationDirection = iota
	DOWN
)

// TODO maybe convert this to something like a builder
func DoGenericMigration(fileSystem fs.FS, dbName string, db *sql.DB, dbType DBType, direction MigrationDirection) error {
	m, err := GetMigrator(fileSystem, dbName, db, dbType)

	if err != nil {
		return err
	}

	if direction == DOWN {
		return m.Down()
	}
	return m.Up()
}

func GetMigrator(fileSystem fs.FS, dbName string, db *sql.DB, dbType DBType) (*migrate.Migrate, error) {
	d, err := iofs.New(fileSystem, migrationFilesPath)

	if err != nil {
		return nil, err
	}

	var dri database.Driver

	if dbType == POSTGRESS {
		dri, _ = postgres.WithInstance(db, &postgres.Config{})
	} else if dbType == SQLITE3 {
		dri, _ = sqlite3.WithInstance(db, &sqlite3.Config{})
	}

	return migrate.NewWithInstance("simpleserer_migrator", d, dbName, dri)
}
