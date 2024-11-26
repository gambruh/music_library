package app

import (
	"context"
	"fmt"

	"github.com/gambruh/music_library/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Storage interface is a data storage. Implementation may vary
type Storage interface {
	GetSong(ctx context.Context, group string, song string) (*storage.Song, error)
	AddSong(ctx context.Context, song *storage.Song) error
	DeleteSong(ctx context.Context, group string, song string) error
	EditSong(ctx context.Context, song *storage.Song) error
}

type Config interface {
	GetConfig() error
	GetAddr() string
}

// TODO
// Logger defines a generic interface for logging.
// type Logger interface {
// 	Debug(msg string, fields map[string]interface{})
// 	Info(msg string, fields map[string]interface{})
// 	Warn(msg string, fields map[string]interface{})
// 	Error(msg string, fields map[string]interface{})
// 	Fatal(msg string, fields map[string]interface{})
// }

// Service struct is used to implement service functions
// ideally, I would implement logger and config as interfaces, to pass into the constructor
// TODO however, for the test task I will skip this
type Service struct {
	Logger  *logrus.Logger
	Config  Config
	Storage Storage
}

// Create a new service with the logger attached to it
func NewService(logger *logrus.Logger, config Config, storage Storage) *Service {
	return &Service{Logger: logger, Config: config, Storage: storage}
}

// handlers for a server
func (s *Service) InitRouter() *echo.Echo {

	// create a new Mux
	r := echo.New()

	// r.Use(middleware.Compress(5, "text/plain", "text/html", "application/json"))

	// Add routes
	s.addRoutes(r)
	return r
}

// Starts a new service on the configured address (config.Address)
func (s *Service) Start() error {
	r := s.InitRouter()
	address := s.Config.GetAddr()
	fmt.Println("get address:", address)
	return r.Start(address)
}
