package main

import (
	"easyvpn/src/database"
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

	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)
	r.HandleFunc("/user/sign-in", routes.UserLogin).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/check-token", routes.CheckUserToken).Methods(http.MethodPost, http.MethodOptions)

	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.CorsMiddleware, middleware.CheckAdminRoute)
	adminRouter.HandleFunc("/user", routes.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	adminRouter.HandleFunc("/user", routes.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	adminRouter.HandleFunc("/user", routes.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	adminRouter.HandleFunc("/user", routes.DeleteUser).Methods(http.MethodPut, http.MethodOptions)

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err = http.ListenAndServe(":"+port, r)

	if err != nil {
		panic("Error starting REST server")
	}
}
