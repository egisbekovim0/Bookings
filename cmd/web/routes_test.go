package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/yerlan/bookings/internal/config"
)
func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		// nothing
	default:
		t.Errorf("type is not chi mux but %T", v)
	}
}