package main

import (
	"easyvpn/src/database"
	"easyvpn/src/groups"
	"easyvpn/src/settings"
	"easyvpn/src/user"
	"easyvpn/src/vpn"
	"net/http"
	"strconv"
	"time"

	"easyvpn/src/utils"
	"embed"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/provider"
	"github.com/go-pkgz/auth/token"
)

var LOGFILE_DIR = `C:\Program Files\OpenVPN\log\server-dev.log`
var CONFIGFILE_DIR = `C:\Program Files\OpenVPN\config-auto\server-dev.ovpn`

//go:embed app/*
var svelte embed.FS

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

	options := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) {
			return "secret", nil
		}),
		TokenDuration:  time.Hour * 24,
		CookieDuration: time.Hour * 24,
		Issuer:         "easy-vpn",
		URL:            "http://localhost:8080",
		AvatarStore:    avatar.NewLocalFS("/tmp"),
		JWTCookieName:  "JWT",
		JWTHeaderKey:   "JWT",
		SendJWTHeader:  true,
		ClaimsUpd: token.ClaimsUpdFunc(func(claims token.Claims) token.Claims {
			if claims.User != nil {
				user, err := user.GetUser(claims.User.Name)
				if err == nil {
					claims.User.SetAdmin(user.IsAdmin)
					claims.User.SetStrAttr("password_expiry", user.PasswordExpiry.Format(time.DateTime))
					claims.User.SetStrAttr("id", strconv.Itoa(int(user.ID)))
					claims.User.SetBoolAttr("enabled", user.Enabled)
				}
			}
			return claims
		}),
		Validator: token.ValidatorFunc(func(_ string, claims token.Claims) bool {
			return claims.User != nil
		}),
	}

	service := auth.NewService(options)
	r := setupRouter(service)

	fmt.Print("Startup Successful")
	err = http.ListenAndServe(fmt.Sprintf(":%d", 8080), r)
	if err != nil {
		panic(err)
	}

}

func setupRouter(service *auth.Service) *chi.Mux {
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

	authRoutes, avaRoutes := service.Handlers()
	r.Mount("/auth", authRoutes)
	r.Mount("/avatar", avaRoutes)

	service.AddDirectProvider("local", provider.CredCheckerFunc(func(username, password string) (ok bool, err error) {
		return user.AuthUser(username, password)
	}))

	m := service.Middleware()

	r.Route("/settings", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", settings.GetSettingsEndpoint)
		r.Put("/", settings.SetSettingsEndpoint)
		r.Get("/file", settings.GetConfigFileEndpoint)
	})

	r.Route("/user", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", user.GetUsersEndpoint)
		r.Post("/", user.CreateUserEndpoint)
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", user.DeleteUserEndpoint)
			r.Put("/", user.UpdateUserEndpoint)
			r.Post("/set-pw", user.SetPWEndpoint)
		})
	})

	r.HandleFunc("/vpn/log", vpn.GetVpnLogs)

	r.Route("/vpn", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", vpn.GetServerStatusEndpoint)
		r.Get("/connections", vpn.GetActiveConnectionsEndpoint)
		r.Post("/", vpn.VpnOperationEndpoint)
	})

	r.Route("/group", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", groups.GetGroupsEndpoint)
		r.Post("/", groups.CreateGroupEndpoint)
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", groups.DeleteGroupEndpoint)
			r.Put("/", groups.UpdateGroupEndpoint)
		})
		r.Route("/membership/{id}", func(r chi.Router) {
			r.Get("/", groups.GetGroupMembershipEndpoint)
			r.Post("/", groups.CreateGroupMembershipEndpoint)
			r.Delete("/", groups.DeleteGroupMembershipEndpoint)
		})
	})
	r.Handle("/app/*", http.FileServer(http.FS(svelte)))
	return r
}
