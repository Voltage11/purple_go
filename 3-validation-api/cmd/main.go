package main

import (
	"net/http"
	"purple2/3-validation-api/configs"
	"purple2/3-validation-api/internal/verify"
)

func main() {
	cfg := configs.NewConfig()

	router := http.NewServeMux()
	verify.NewVarifyHandler(router, cfg)
	

	server := &http.Server{
		Addr: ":8081",
		Handler: router,
	}

	server.ListenAndServe()

}