package router

import (
	"log"
	"net/http"
	"strings"

	"github.com/srazap/luckbuy/internal"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ep := strings.Split(r.URL.Path, "/")
		if protected, ok := protectedRoutes[ep[len(ep)-1]]; ok {
			if protected {
				sessionId := r.Header.Get("session_id")
				if sessionId != "" {
					// get user by session id
					ok, err := internal.ValidateSession(sessionId)
					if !ok || err != nil {
						if err != nil {
							log.Print(err)
						}
						http.Error(w, "invalid or expired session id", http.StatusForbidden)
						return
					}
				} else {
					http.Error(w, "invalid or expired session id", http.StatusForbidden)
					return
				}
			}
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, "Not Found", http.StatusNotFound)
	})
}
