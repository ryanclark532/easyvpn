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
	"strconv"
	"time"

	"easyvpn/src/utils"
	"embed"
	"fmt"

	"github.com/a-h/templ"
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

	r.Route("/user", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Route("/{id}", func(r chi.Router) {
			r.Post("/set-pw", user.SetPWEndpoint)
		})
	})

	r.Route("/user/config/{username}", func(r chi.Router) {
		r.Use(m.Auth)
		r.Post("/", user.CreateUserConfigEndpoint)
		r.Delete("/", user.DeleteUserConfigEndpoint)
	})

	r.Route("/group", func(r chi.Router) {
		r.Use(m.AdminOnly)
		r.Get("/", groups.GetGroupsEndpoint)
		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", groups.UpdateGroupEndpoint)
		})
		r.Route("/membership/{id}", func(r chi.Router) {
			r.Get("/", groups.GetGroupMembershipEndpoint)
			r.Post("/", groups.CreateGroupMembershipEndpoint)
			r.Delete("/", groups.DeleteGroupMembershipEndpoint)
		})
	})

	r.Handle("/static/*", http.FileServer(http.FS(static)))

	r.Handle("/", templ.Handler(common.Home()))

	r.Route("/login", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			login.Login().Render(r.Context(), w)
		})
		r.Post("/", login.HandleLogin)
	})
	r.Route("/users", func(r chi.Router) {
		r.Use(login.AuthMiddleware)
		r.Get("/", user.UsersPage)
		r.Post("/", user.CreateNewUser)
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", user.DeleteUser)
			r.Put("/", user.UpdateUser)
		})
	})

	r.Route("/groups", func(r chi.Router) {
		r.Use(login.AuthMiddleware)
		r.Get("/", groups.GroupsPage)
		r.Post("/", groups.CreateGroup)
		r.Route("/{id}", func(r chi.Router) {
			r.Delete("/", groups.DeleteGroup)
		})
	})

	r.Route("/settings", func(r chi.Router) {
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
		r.Get("/active-connections", vpn.GetActiveUsersPage)
		r.Get("/logs", vpn.GetVpnLogsPage)
		r.Post("/disconnect", vpn.DisconnectClient)
	})

	return r
}
