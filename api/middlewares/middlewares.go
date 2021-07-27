package middlewares

import (
	"log"
	"microservicios/api/restutil"
	"microservicios/security"
	"net/http"
	"time"
)

func LogRequests(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		next(w, r)
		log.Printf(`{"proto": "%s", "method": "%s", "route": "%s%s", "request_time": "%v"}`,
			r.Proto, r.Method, r.Host, r.URL.Path, time.Since(t))
	}
}

func Authenticare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := security.ExtractToken(r)
		if err != nil {
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnautorized)
			return
		}

		token, err := security.ParseToken(tokenString)
		if err != nil {
			log.Println("invalid token:", tokenString)
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnautorized)
			return
		}

		if !token.Valid {
			log.Println("error on parse token:", err.Error())
			restutil.WriteError(w, http.StatusUnauthorized, restutil.ErrUnautorized)
			return
		}

		next(w, r)
	}
}
