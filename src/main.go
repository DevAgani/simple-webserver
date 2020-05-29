package main

import (
	"log"
	"net/http"
)

type server struct{}

// ServeHTTP method on the server struct is implementing the Handler interface
// we're setting the content type to application/json and the return a ok status
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
