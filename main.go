package main

import (
	"fmt"
	"log"
	"net/http"
);

func main()  {
	fmt.Print("Savepoint API is starting...");
	runServer();
}

func runServer() {
	//listen for request on path /bar
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to SavePoint API")
	})

	// start server
 	log.Fatal(http.ListenAndServe(":8080", nil));
}

