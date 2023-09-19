package main

import (
	"easyvpn/src/database"
	"easyvpn/src/utils"
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

	err = utils.SetupVPNServer()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)
	r.HandleFunc("/user/sign-in", routes.UserLogin).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/user/check-token", routes.CheckUserToken).Methods(http.MethodPost, http.MethodOptions)

	adminRouter := r.PathPrefix("/").Subrouter()
	adminRouter.Use(middleware.CorsMiddleware, middleware.CheckAdminRoute)
	adminRouter.HandleFunc("/user", routes.GetUsers).Methods(http.MethodGet, http.MethodOptions)
	adminRouter.HandleFunc("/user", routes.CreateUser).Methods(http.MethodPost, http.MethodOptions)
	adminRouter.HandleFunc("/user", routes.DeleteUser).Methods(http.MethodDelete, http.MethodOptions)
	adminRouter.HandleFunc("/user", routes.UpdateUser).Methods(http.MethodPut, http.MethodOptions)
	adminRouter.HandleFunc("/user/set-temporary-password", routes.SetTemporaryPassword).Methods(http.MethodPut, http.MethodOptions)

	userRouter := r.PathPrefix("/").Subrouter()
	userRouter.Use(middleware.CorsMiddleware, middleware.CheckUserRoute)
	userRouter.HandleFunc("/user/change-password", routes.ChangeUserPassword).Methods(http.MethodPost, http.MethodOptions)

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err = http.ListenAndServe(":"+port, r)

	if err != nil {
		panic("Error starting REST server")
	}
}
