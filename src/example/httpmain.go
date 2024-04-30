package main

import (
	"log"
	"myHttpServer/httpServer"
	"net/http"
)

func main() {
	// fmt.Println("success")
	handler := http.HandlerFunc(httpServer.PalyerServer)

	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
