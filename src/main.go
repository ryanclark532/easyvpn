package main

import (
	"easyvpn/src/auth"
	"easyvpn/src/database"
	"easyvpn/src/middleware"
	"easyvpn/src/user"
	"easyvpn/src/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	db := make(chan error)
	vpn := make(chan error)

	go func() {
		err := database.InitializeDatabase()
		db <- err
	}()

	go func() {
		err := utils.SetupVPNServer()
		vpn <- err

		err = utils.StartVPNServer()
		vpn <- err

	}()

	dberr := <-db
	vpnerr := <-vpn

	if dberr != nil {
		panic(dberr)
	}

	if vpnerr != nil {
		panic(vpnerr)
	}

	r := SetupRouter()

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		panic("Error starting REST server")
	}
}

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.CorsMiddleware)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./dist/static")))) // TODO make this configurable per dev vs prod
	r.HandleFunc("/auth/sign-in", auth.UserLoginEndpoint).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/auth/check-token", auth.CheckUserTokenEndpoint).Methods(http.MethodPost, http.MethodOptions)

	adminRouter := r.PathPrefix("/").Subrouter()
	adminRouter.Use(middleware.CorsMiddleware, middleware.CheckAdminRoute)
	adminRouter.HandleFunc("/user", user.GetUsersEndpoint).Methods(http.MethodGet, http.MethodOptions)
	adminRouter.HandleFunc("/user", user.CreateUserEndpoint).Methods(http.MethodPost, http.MethodOptions)
	adminRouter.HandleFunc("/user", user.DeleteUserEndpoint).Methods(http.MethodDelete, http.MethodOptions)
	adminRouter.HandleFunc("/user", user.UpdateUserEndpoint).Methods(http.MethodPut, http.MethodOptions)
	adminRouter.HandleFunc("/auth/set-temporary-password", auth.SetTemporaryPasswordEndpoint).Methods(http.MethodPut, http.MethodOptions)

	userRouter := r.PathPrefix("/").Subrouter()
	userRouter.Use(middleware.CorsMiddleware, middleware.CheckUserRoute)
	userRouter.HandleFunc("/auth/change-password", auth.ChangeUserPasswordEndpoint).Methods(http.MethodPost, http.MethodOptions)

	return r
}
