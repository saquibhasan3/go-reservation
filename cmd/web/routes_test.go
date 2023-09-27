package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/saquibhasan3/go-reservation/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// Do nothing got what expected
	default:
		t.Error(fmt.Sprintf("Expecting chi mux but got %T ", v))

	}
}
