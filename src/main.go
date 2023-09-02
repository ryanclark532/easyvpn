package main

import (
	"easyvpn/src/database"
	"easyvpn/src/services"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"easyvpn/src/middleware"
	"easyvpn/src/routes"
)

func main() {
	err := database.InitializeDatabase()
	if err != nil {
		panic(err)
	}

	users, _ := services.GetUsers()
	fmt.Println(users)

	r := mux.NewRouter()
	r.HandleFunc("/user/sign-in", routes.UserLogin).Methods(http.MethodPost)
	r.HandleFunc("/user/check-token", routes.CheckUserToken).Methods(http.MethodPost)
	r.HandleFunc("/user", routes.CreateUser).Methods(http.MethodPost)
	corsHandler := middleware.SetupCORS(r)

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err = http.ListenAndServe(":"+port, corsHandler)
	if err != nil {
		panic("Error starting REST server")
	}
}
