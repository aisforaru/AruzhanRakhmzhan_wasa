package main

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(
		handlers.AllowedOrigins([]string{
			"http://localhost:5173",
		}),
		handlers.AllowedMethods([]string{
			"GET", "POST", "OPTIONS", "DELETE", "PUT",
		}),
		handlers.AllowedHeaders([]string{
			"Content-Type",
			"Authorization",
		}),
		handlers.AllowCredentials(),
	)(h)
}


