// Package tools +build tools
package tools

import _ "github.com/99designs/gqlgen"
import _ "github.com/nrfta/go-graphql-scalars"

//go:generate go run github.com/99designs/gqlgen generate
