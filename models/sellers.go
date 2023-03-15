package models

import (
	"encoding/json"
	"github.com/Vladimir1k/OnlineShop/repository"
	"log"
	"net/http"
	"strconv"
)

type Seller struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func FindSellers(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "GET" {
		w.Write([]byte("choose correct method"))
		return
	}

	stmt, err := db.Query("SELECT * FROM sellers")
	if err != nil {
		log.Println(err)
	}

	var sellers []Seller

	for stmt.Next() {
		var sel Seller
		err = stmt.Scan(&sel.ID, &sel.Name, &sel.Phone)
		if err != nil {
			log.Println(err)
		}
		sellers = append(sellers, sel)
	}
	sellersBytes, _ := json.MarshalIndent(sellers, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(sellersBytes)

	defer stmt.Close()
	defer db.Close()
}

func FindSellerById(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "GET" {
		w.Write([]byte("choose correct method"))
		return
	}

	emp, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || emp < 1 {
		http.NotFound(w, r)
		return
	}

	stmt := `SELECT * FROM sellers WHERE id=$1`
	if err != nil {
		log.Println(err)
	}

	row := db.QueryRow(stmt, emp)

	var seller Seller
	err = row.Scan(&seller.ID, &seller.Name, &seller.Phone)
	if err != nil {
		log.Println(err)
	}

	sellerBytes, _ := json.MarshalIndent(seller, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(sellerBytes)

	defer db.Close()
}

func EditSeller(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "PUT" {
		w.Write([]byte("choose correct method"))
		return
	}

	emp, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || emp < 1 {
		http.NotFound(w, r)
		return
	}

	var seller Seller

	err = json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt := `UPDATE sellers SET name=$1, phone=$2 WHERE id=$3`
	if err != nil {
		log.Println(err)
	}

	db.QueryRow(stmt, seller.Name, seller.Phone, emp)

	w.Write([]byte("Update"))

	defer db.Close()
}

func AddSeller(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "POST" {
		w.Write([]byte("choose correct method"))
		return
	}

	var seller Seller

	err = json.NewDecoder(r.Body).Decode(&seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO sellers (name, phone) VALUES ($1, $2)")
	stmt.Exec(seller.Name, seller.Phone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("bad requset:", err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
	defer stmt.Close()
}

func RemoveSeller(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "DELETE" {
		w.Write([]byte("choose correct method"))
		return
	}

	emp, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || emp < 1 {
		http.NotFound(w, r)
		return
	}

	stmt := `DELETE FROM sellers WHERE id=$1`
	if err != nil {
		log.Println(err)
	}

	db.QueryRow(stmt, emp)
	w.Write([]byte("deleted"))

	defer db.Close()
}
