package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Storage interface is a data storage. Implementation may vary
type Storage interface {
	GetSong(group, song string) error
	AddSong(group, song string) error
	DeleteSong(group, song string) error
	EditSong(group, song string) error
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
	logger  *logrus.Logger
	config  Config
	storage Storage
}

// Create a new service with the logger attached to it
func NewService(logger *logrus.Logger, config Config, storage Storage) *Service {
	return &Service{logger: logger, config: config, storage: storage}
}

// handlers for a server
func (h *Service) InitRouter() *echo.Echo {

	// create a new Mux
	r := echo.New()

	// r.Use(middleware.Compress(5, "text/plain", "text/html", "application/json"))

	// Add routes
	addRoutes(r)
	return r
}

// Starts a new service on the configured address (config.Address)
func (h *Service) Start() error {
	r := h.InitRouter()
	address := h.config.GetAddr()
	return r.Start(address)
}
