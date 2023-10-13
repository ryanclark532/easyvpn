package main

import (
	"easyvpn/src/database"
	"easyvpn/src/groups"
	"easyvpn/src/user"
	"easyvpn/src/vpn"
	"net/http"
	"strings"
	"time"

	"easyvpn/src/utils"
	"embed"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
)

//go:embed app/*
var svelte embed.FS

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

	options := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return "secret", nil
		}),
		TokenDuration:  time.Minute * 5, // token expires in 5 minutes
		CookieDuration: time.Hour * 24,  // cookie expires in 1 day and will enforce re-login
		Issuer:         "easy-vpn",
		URL:            "http://localhost:8080",
		AvatarStore:    avatar.NewLocalFS("/tmp"),
		Validator: token.ValidatorFunc(func(_ string, claims token.Claims) bool {
			// allow only dev_* names
			return claims.User != nil && strings.HasPrefix(claims.User.Name, "dev_")
		}),
	}

	service := auth.NewService(options)

	r := setupRouter(service)

	port := "8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		panic(err)
	}
}

/*
	func SetupRouter() *mux.Router {
		r := mux.NewRouter()
		r.Use(middleware.CorsMiddleware)

		apiRouter := r.PathPrefix("/api").Subrouter()
		apiRouter.HandleFunc("/auth/sign-in", auth.UserLoginEndpoint).Methods(http.MethodPost, http.MethodOptions)
		apiRouter.HandleFunc("/auth/check-token", auth.CheckUserTokenEndpoint).Methods(http.MethodPost, http.MethodOptions)

		adminRouter := apiRouter.PathPrefix("/").Subrouter()
		adminRouter.Use(middleware.CorsMiddleware, middleware.CheckAdminRoute)
		adminRouter.HandleFunc("/user", user.GetUsersEndpoint).Methods(http.MethodGet, http.MethodOptions)
		adminRouter.HandleFunc("/user", user.CreateUserEndpoint).Methods(http.MethodPost, http.MethodOptions)
		adminRouter.HandleFunc("/user", user.DeleteUserEndpoint).Methods(http.MethodDelete, http.MethodOptions)
		adminRouter.HandleFunc("/user", user.UpdateUserEndpoint).Methods(http.MethodPut, http.MethodOptions)
		adminRouter.HandleFunc("/auth/set-temporary-password", auth.SetTemporaryPasswordEndpoint).Methods(http.MethodPut, http.MethodOptions)
		adminRouter.HandleFunc("/vpn", vpn.GetServerStatusEndpoint).Methods(http.MethodGet, http.MethodOptions)
		adminRouter.HandleFunc("/vpn/operation", vpn.VpnOperationEndpoint).Methods(http.MethodPost, http.MethodOptions)
		adminRouter.HandleFunc("/vpn/connections", vpn.GetActiveConnectionsEndpoint).Methods(http.MethodGet, http.MethodOptions)
		adminRouter.HandleFunc("/groups", groups.GetGroupsEndpoint).Methods(http.MethodOptions, http.MethodGet)

		userRouter := apiRouter.PathPrefix("/").Subrouter()
		userRouter.Use(middleware.CorsMiddleware, middleware.CheckUserRoute)
		userRouter.HandleFunc("/auth/change-password", auth.ChangeUserPasswordEndpoint).Methods(http.MethodPost, http.MethodOptions)
		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/app", http.StatusSeeOther)
		})
		r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.FS(svelte))))

		return r
	}
*/
func setupRouter(service *auth.Service) *chi.Mux {
	m := service.Middleware()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", user.GetUsersEndpoint)
		r.Post("/", user.CreateUserEndpoint)
		r.Delete("/", user.DeleteUserEndpoint)
		r.Put("/", user.UpdateUserEndpoint)
	})

	r.Route("/vpn", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", vpn.GetServerStatusEndpoint)
		r.Get("/connections", vpn.GetActiveConnectionsEndpoint)
		r.Post("/", vpn.VpnOperationEndpoint)
	})

	r.Route("/group", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", groups.GetGroupsEndpoint)
		r.Get("/memberships", groups.GetGroupMembershipEndpoint)
	})

	authRoutes, avaRoutes := service.Handlers()
	r.Mount("/auth", authRoutes)
	r.Mount("/avatar", avaRoutes)

	service.AddDirectProvider("local", provider.CredCheckerFunc(func(user, password string) (ok bool, err error) {
		return true, err
	}))

	return r
}
