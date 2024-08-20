package main

import (
	"fmt"
	"net/http"

	"github.com/avukadin/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API Service...")
	//cooler GO API println
	err := http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Error(err)
	}

}
