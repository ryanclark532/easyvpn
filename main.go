package main

import (
	"easyvpn/src/common"
	"easyvpn/src/database"
	"easyvpn/src/groups"
	"easyvpn/src/login"
	"easyvpn/src/settings"
	"easyvpn/src/user"
	"easyvpn/src/vpn"
	"net/http"

	"easyvpn/src/utils"
	"embed"
	"fmt"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

//go:embed static/*
var static embed.FS

func main() {

	err := database.Test()
	if err != nil {
		panic(err)
	}

	vpn := make(chan error)

	go func() {
		err := utils.SetupVPNServer()
		vpn <- err
	}()

	vpnerr := <-vpn

	if vpnerr != nil {
		panic(vpnerr)
	}

	r := setupRouter()

	fmt.Print("Startup Successful")
	err = http.ListenAndServe(fmt.Sprintf(":%d", 8080), r)
	if err != nil {
		panic(err)
	}

}

func setupRouter() *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*"}, //TO Be changes in real deployment
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "Set-Cookie", "Jwt"},
		ExposedHeaders:   []string{"Jwt"},
		AllowCredentials: true,
	})

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.FileServer(http.FS(static)))

	r.Handle("/", templ.Handler(common.Home()))

	r.Route("/login", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			login.Login().Render(r.Context(), w)
		})
		r.Post("/", login.HandleLogin)
		r.Post("/signout", login.HandleSignout)
	})
	r.Route("/users", func(r chi.Router) {
		r.Use(login.AuthMiddleware)
		r.Get("/", user.UsersPage)
		r.Post("/", user.CreateNewUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", user.DeleteUser)
			r.Post("/", user.UpdateUser)
		})
	})

	r.Route("/groups", func(r chi.Router) {
		r.Use(login.AuthMiddleware)
		r.Get("/", groups.GroupsPage)
		r.Post("/", groups.CreateGroup)
		r.Route("/{id}", func(r chi.Router) {
			r.Post("/", groups.UpdateGroupPage)
			r.Delete("/", groups.DeleteGroup)
		})
	})

	r.Route("/settings", func(r chi.Router) {
		r.Use(login.AuthMiddleware)
		r.Route("/server", func(r chi.Router) {
			r.Get("/", settings.ServerSettingsPage)
			r.Post("/", settings.SetServerSettings)
		})

		r.Route("/client", func(r chi.Router) {
			r.Get("/", settings.ClientSettingsPage)
			r.Post("/", settings.SetClientSettings)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Get("/", settings.AuthSettingsPage)
			r.Post("/", settings.SetAuthSettings)
		})

		r.Route("/config", func(r chi.Router) {
			r.Get("/", settings.ConfigFileSettingsPage)
			r.Post("/", settings.SetConfigFileContent)
		})
	})

	r.Route("/vpn", func(r chi.Router) {
		r.Use(login.AuthMiddleware)
		r.Get("/active-connections", vpn.GetActiveUsersPage)
		r.Get("/logs", vpn.GetVpnLogsPage)
		r.Post("/disconnect", vpn.DisconnectClient)
		r.Get("/overview", vpn.StatusOverviewPage)
		r.Post("/", vpn.VpnOperation)
	})

	return r
}
