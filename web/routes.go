package web

import (
	"log"
	"net/http"
)

func (s Server) Route() {
	http.HandleFunc("/", s.SumHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
