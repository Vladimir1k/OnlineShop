package repository

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

type auth struct {
	User     string
	Password string
}

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authData = auth{}

		err := godotenv.Load("../OnlineShop/auth.env")
		if err != nil {
			log.Println("Error loading .env")
		}

		authData.User = os.Getenv("SHOP_USERNAME")
		authData.Password = os.Getenv("SHOP_PASS")

		user, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if authData.User != user || authData.Password != strings.ReplaceAll(password, "\n", "") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		handler.ServeHTTP(w, r)
	}
}
