// Package storage provides Postgres DB implementation of Storage interface
package storage

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Storage interface is used to abstract the actual storage
// it would be possible to use some inmemory structure for tests
// or provide anything else
type Storage interface {
	GetSong(group, song string) error
	AddSong(group, song string) error
	DeleteSong(group, song string) error
	EditSong(group, song string) error
}

type SQLdb struct {
	DB *sql.DB
}

// NewSQLdb() returns an open connection to the database
func NewSQLdb(postgresStr string) (*SQLdb, error) {
	DB, err := sql.Open("pgx", postgresStr)
	if err != nil {
		return nil, err
	}

	return &SQLdb{
		DB: DB,
	}, nil
}

// GetDB() is a wrapper for Storage interface
func GetDB(postgresStr string) (*SQLdb, error) {

	db, err := NewSQLdb(postgresStr)
	if err != nil {
		return nil, err
	}

	err = db.InitDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *SQLdb) CheckConn(dbAddress string) error {
	db, err := sql.Open("pgx", dbAddress)
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

// InitDatabase creates tables in an empty database
func (s *SQLdb) InitDatabase() error {
	dbURL := os.Getenv("MUSIC_DATABASE_STRING")
	if dbURL == "" {
		return errors.New("env variable MUSIC_DATABASE_STRING is not set")
	}

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		// s.logger.Errorf("Failed to get working directory: %v ", err)
		return err
	}
	migrationPath := "file://" + filepath.Join(wd, "internal", "storage", "database", "migrations")

	m, err := migrate.New(
		migrationPath, // Folder with migration files
		dbURL,         // connection string for db
	)
	if err != nil {
		// s.logger.Errorf("Failed to initialize migrations: %v", err)
		return err
	}

	// s.logger.Info("Starting migrations...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		// s.logger.Errorf("Failed to apply migrations: %v", err)
		return err
	}
	// s.logger.Info("Migrations applied successfully!")

	return nil
}

func (s *SQLdb) AddSong(group, song string) error {
	return nil
}

func (s *SQLdb) GetSong(group, song string) error {
	return nil
}

func (s *SQLdb) DeleteSong(group, song string) error {
	return nil
}

func (s *SQLdb) EditSong(group, song string) error {
	return nil
}
