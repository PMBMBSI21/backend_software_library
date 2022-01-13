package middlewares

import (
	"errors"
	"net/http"
	"software_library/backend/api/auth"
	"software_library/backend/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// semua origin mendapat ijin akses
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// // w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000/login")

		// // semua method diperbolehkan masuk
		// w.Header().Set("Access-Control-Allow-Methods", "*")

		// // semua header diperbolehkan untuk disisipkan
		// w.Header().Set("Access-Control-Allow-Headers", "*")

		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthenticationAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		cekAdmin, _ := auth.CekAdmin(r)
		if cekAdmin != 1 {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("you dont have access"))
			return
		}
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
