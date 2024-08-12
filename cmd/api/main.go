package main

import (
	"fmt"
	"net/http"

	"github.com/LinPaing21/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go Api Service...")

	fmt.Println(`
 _____ _____    _____ _____ _ 
|   __|     |  |  _  |  _  |_|
|  |  |  |  |  |     |   __| |
|_____|_____|  |__|__|__|  |_|`)

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
