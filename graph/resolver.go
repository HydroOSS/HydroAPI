//go:generate go run github.com/99designs/gqlgen

package graph

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the main resolver used by the API.
type Resolver struct {
	DB *r.Session
}
