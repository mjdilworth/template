package server

import (
	"net/http"
)

func (s server) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// Write response status code
	w.WriteHeader(http.StatusOK)
	//fmt.Fprintf(w, `{"message": "I am healthy"}`+"\n")
	w.Write([]byte(`{"message": "I am healthy"}`))

}
