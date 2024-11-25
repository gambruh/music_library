// This is a CRUD service that provides an API to:
// - add song to the database
// - get song from the database
// - edit song details in the database
// - delete song from the database

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/gambruh/music_library/internal/app"
	"github.com/gambruh/music_library/internal/config"
	"github.com/gambruh/music_library/internal/logging"
	storage "github.com/gambruh/music_library/internal/storage/database"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// create a new config struct
	cfg := config.NewConfig()
	err := cfg.GetConfig()
	if err != nil {
		return err
	}

	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		// s.logger.Errorf("Failed to get working directory: %v ", err)
		return err
	}
	logFilePath := filepath.Join(wd, cfg.GetLogFile())

	// create log file, start logger
	logger, err := logging.NewLogger(logFilePath)
	if err != nil {
		return err
	}

	// initialize connection to the database
	db, err := storage.GetDB(cfg.GetDatabaseConnStr())
	if err != nil {
		return err
	}

	// init a music library server
	s := app.NewService(logger, cfg, db)

	// Use a goroutine to start the music library server
	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(w, "Failed to start server: %v\n", err)
			cancel() // Cancel the context to stop the app
		}
	}()

	// Wait for the context to be canceled (on interrupt)
	<-ctx.Done()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
