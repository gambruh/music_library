// Package storage provides Postgres DB implementation of Storage interface
package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gambruh/music_library/internal/storage"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type SQLdb struct {
	DB *sql.DB
}

// NewSQLdb() returns an open connection to the database
func NewSQLdb(postgresStr string) (*SQLdb, error) {
	fmt.Println("Connecting to the database:", postgresStr)
	DB, err := sql.Open("pgx", postgresStr)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database")
	return &SQLdb{
		DB: DB,
	}, nil
}

// function GetDB() returns database which implements Storage interface
func GetDB() (*SQLdb, error) {
	dbURL := os.Getenv("MUSIC_DATABASE_STRING")

	db, err := NewSQLdb(dbURL)
	if err != nil {
		return nil, err
	}

	err = db.InitDatabase()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// function to ping the db
func (s *SQLdb) CheckConn() error {
	dbURL := os.Getenv("MUSIC_DATABASE_STRING")

	// dbURL = "user=postgres password=postgres host=postgres port=5432 dbname=music_library sslmode=disable"

	fmt.Println("FROM CHECKCONNL:", dbURL)
	db, err := sql.Open("pgx", dbURL)
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
	dbURL := os.Getenv("MUSIC_DATABASE_URL")
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

func (s *SQLdb) AddSong(ctx context.Context, song *storage.Song) error {
	// Get query parameters
	group := song.Group
	song_name := song.Name
	lyrics := song.Text
	link := song.Link
	releaseDate := song.ReleaseDate

	// Execute the add song query
	_, err := s.DB.ExecContext(ctx, ADD_SONG_QUERY, group, song_name, releaseDate, lyrics, link)
	if err != nil {
		return fmt.Errorf("failed to add song: %w", err)
	}

	return nil
}

func (s *SQLdb) GetSong(ctx context.Context, groupName, songName string) (*storage.Song, error) {
	var song storage.Song

	// Execute the get song query
	err := s.DB.QueryRowContext(ctx, GET_SONG_QUERY, groupName, songName).Scan(
		&song.Name, &song.Group, &song.ReleaseDate, &song.Text, &song.Link,
	)

	if err != nil {
		return nil, err
	}
	return &song, nil
}

func (s *SQLdb) DeleteSong(ctx context.Context, group, song string) error {

	// Execute the delete query
	result, err := s.DB.ExecContext(ctx, DEL_SONG_QUERY, group, song)
	if err != nil {
		return fmt.Errorf("failed to delete song: %w", err)
	}

	// Check the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no song found with group %q and title %q", group, song)
	}

	return nil

}

func (s *SQLdb) EditSong(ctx context.Context, song *storage.Song) error {
	// Get query parameters
	group := song.Group
	song_name := song.Name
	releaseDate := song.ReleaseDate
	lyrics := song.Text
	link := song.Link

	if releaseDate != "" {

	}

	// Execute the update query
	result, err := s.DB.ExecContext(ctx, EDIT_SONG_QUERY, group, song_name, releaseDate, lyrics, link)
	if err != nil {
		return fmt.Errorf("failed to edit song: %w", err)
	}

	// Check the number of affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no song found with group %q and title %q", group, song)
	}

	return nil
}
