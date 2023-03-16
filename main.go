package main

import (
	"github.com/Vladimir1k/OnlineShop/models"
	"github.com/Vladimir1k/OnlineShop/repository"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/sellers", repository.BasicAuth(models.FindSellers))
	mux.HandleFunc("/remove/seller", repository.BasicAuth(models.RemoveSeller))
	mux.HandleFunc("/edit/seller", repository.BasicAuth(models.EditSeller))
	mux.HandleFunc("/add/sellers", repository.BasicAuth(models.AddSeller))
	mux.HandleFunc("/find/seller", repository.BasicAuth(models.FindSellerById))

	mux.HandleFunc("/buyers", repository.BasicAuth(models.FindBuyers))
	mux.HandleFunc("/remove/buyers", repository.BasicAuth(models.RemoveBuyer))
	mux.HandleFunc("/edit/buyers", repository.BasicAuth(models.EditBuyer))
	mux.HandleFunc("/add/buyers", repository.BasicAuth(models.AddBuyer))
	mux.HandleFunc("/find/buyer", repository.BasicAuth(models.FindBuyerById))

	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func Concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func Concat2(str []string) string {
	var sb strings.Builder
	for _, v := range str {
		sb.WriteString(v)
	}
	return sb.String()
}
