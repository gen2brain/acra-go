package database

import (
	"github.com/gen2brain/acra-go/acra"
)

// DB interface
type DB interface {
	// Open opens database
	Open(file string) error

	// Put puts report
	Put(key string, value acra.Report) error

	// Get gets report
	Get(key string) (acra.Report, error)

	// GetAll gets all reports
	GetAll() ([]acra.Report, error)

	// Delete deletes report
	Delete(key string) error

	// Close closes database
	Close() error
}
