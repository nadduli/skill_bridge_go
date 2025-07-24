package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
	mux := http.NewServeMux()

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)

	server := &http.Server{
		Addr: serverAddr,
		Handler: nil,
	}

	if err := server.ListenAndServe(); err != nil{
		log.Fatalf("Failed to start server! %v", err)
	}

}