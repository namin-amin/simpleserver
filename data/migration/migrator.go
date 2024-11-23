package migration

import (
	"database/sql"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/namin-amin/simpleserver/config"
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

type Migrator struct {
	config        *config.Config
	db            *sql.DB
	dbType        DBType
	scriptsFolder fs.FS
	dbName        string
}

// Do a UP/DOWN migration to all the way up to the latest version
func (m *Migrator) DoGenericMigration(direction MigrationDirection) error {
	mg, err := m.GetMigrator()

	if err != nil {
		return err
	}

	if direction == DOWN {
		return mg.Down()
	}
	return mg.Up()
}

// Get the Migrate to customise your migrations yourself
func (m *Migrator) GetMigrator() (*migrate.Migrate, error) {
	c := config.NewConfig()

	overriddentMigrationPath := c.GetEnvVarWithDefault("migrationpath", migrationFilesPath)

	d, err := iofs.New(m.scriptsFolder, overriddentMigrationPath)

	if err != nil {
		return nil, err
	}

	var driver database.Driver

	if m.dbType == POSTGRESS {
		driver, _ = postgres.WithInstance(m.db, &postgres.Config{})
	} else if m.dbType == SQLITE3 {
		driver, _ = sqlite3.WithInstance(m.db, &sqlite3.Config{})
	}

	return migrate.NewWithInstance("simpleserer_migrator", d, m.dbName, driver)
}

// Creates and Returns the Migrator for doing Database Migrations
//
// Internally this uses go-migrate package so check the docs to setup sql files
func NewMigrator(fileSystem fs.FS, dbName string, db *sql.DB, dbType DBType) *Migrator {
	return &Migrator{
		config:        config.NewConfig(),
		db:            db,
		dbType:        dbType,
		scriptsFolder: fileSystem,
		dbName:        dbName,
	}
}
