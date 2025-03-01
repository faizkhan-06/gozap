package routes

import (
	"net/http"

	"github.com/faizkhan-06/gozap/src/handlers"
)

func RegisterRoutes()(*http.ServeMux) {
	router := http.NewServeMux()

	router.HandleFunc("POST /url", handlers.CreateUrlHandler)


	return router
}