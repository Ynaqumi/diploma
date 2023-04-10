package server

import (
	"fmt"
	"log"
	"net/http"
)

import (
	"github.com/gorilla/mux"
)

func Server() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection)

	server := &http.Server{
		Addr:    "localhost:8282",
		Handler: r,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Ошибка запуска сервиса: %v", err)
	}
}

func handleConnection(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "OK")
}
