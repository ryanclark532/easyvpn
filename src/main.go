package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"easyvpn/src/routes"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user/sign-in", routes.UserLogin).Methods(http.MethodPost)

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
