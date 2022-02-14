package router

import (
	"github.com/gorilla/mux"
	"github.com/srazap/luckbuy/api"
)

var protectedRoutes map[string]bool

func init() {
	protectedRoutes = make(map[string]bool)
}

func HandleRoutes(r *mux.Router) *mux.Router {
	sr := r.PathPrefix("/api/v1").Subrouter()

	// use authentication middleware
	sr.Use(AuthMiddleware)

	// unprotected routes
	protectedRoutes["signup"] = false
	sr.HandleFunc("/signup", api.Signup)
	protectedRoutes["login"] = false
	sr.HandleFunc("/login", api.Login)

	// protected routes
	protectedRoutes["logout"] = true
	sr.HandleFunc("/logout", api.Logout)

	protectedRoutes["points"] = true
	sr.HandleFunc("/points", api.MyPoints)

	protectedRoutes["leaderboard"] = true
	sr.HandleFunc("/leaderboard", api.Leaderboard)

	return sr
}
