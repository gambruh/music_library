// Package storage declares Storage interface, as well as types and errors for it
package storage

import (
	"context"
	"errors"
)

// Storage interface is used to abstract the actual storage
// it would be possible to use some inmemory structure for tests
// or provide anything else
type Storage interface {
	GetSong(ctx context.Context, group string, song string) (*Song, error)
	AddSong(ctx context.Context, song *Song) error
	DeleteSong(ctx context.Context, group string, song string) error
	EditSong(ctx context.Context, song *Song) error
}

type Song struct {
	Name        string `json:"song"`
	Group       string `json:"group"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// errors
var (
	ErrTableDoesntExist = errors.New("table doesn't exist")
	ErrDataNotFound     = errors.New("requested data not found in storage")
)
