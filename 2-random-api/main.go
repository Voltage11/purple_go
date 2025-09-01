package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			w.Write([]byte(fmt.Sprintf("%d", rand.Intn(6) + 1)))
		}
	})

	server := &http.Server{
		Addr: ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}