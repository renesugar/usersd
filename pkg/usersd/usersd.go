// Copyright 2018 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the WTFPL.

package usersd

import (
	"github.com/dgraph-io/badger"
)

// DefaultOptions are the commonly used options for a simple Init call.
var DefaultOptions = Options{
	AdminPassword: "admin",
}

// Service is an authentication and authorization service.
type Service struct {
	DB    *badger.DB
	Index Index

	opts Options
	err  error
}

// New creates and starts a service. Receives an Options instance as argument
// and returns a Service instance and an error if any.
func New(opts Options) (*Service, error) {
	s := new(Service)
	s.opts = opts

	if err := s.Start(); err != nil {
		return nil, err
	}

	return s, nil
}

// Start initialize the service (databases, search indexes, etc...). Returns an
// error if any.
func (s *Service) Start() error {
	return s.openDB()
}

// Close terminates the service (databases, search indexes, etc...). Any error
// closing the service will be stored at Service.err and will be accessible
// from Service.Err().
func (s *Service) Close() {
	s.err = s.closeDB()
}

// Err checks if any error occurred during some processes (closing, etc...).
func (s *Service) Err() error {
	return s.err
}

// IsTemp returns true if the service persistent storage is temporary.
func (s *Service) IsTemp() bool {
	if s.opts.Database == "" {
		return true
	}

	return false
}

// Options are parameters for initializing a service.
type Options struct {
	// Administrator password.
	AdminPassword string

	// Database location, if an empty string is given, a temporary storage will
	// be used.
	Database string
}
