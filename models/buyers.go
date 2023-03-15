package models

import (
	"encoding/json"
	"github.com/Vladimir1k/OnlineShop/repository"
	"log"
	"net/http"
	"strconv"
)

type Buyer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func FindBuyers(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "GET" {
		w.Write([]byte("choose correct method"))
		return
	}

	stmt, err := db.Query("SELECT * FROM buyers")
	if err != nil {
		log.Println(err)
	}

	var buyers []Buyer

	for stmt.Next() {
		var buyer Buyer
		err = stmt.Scan(&buyer.ID, &buyer.Name, &buyer.Phone, &buyer.Address)
		if err != nil {
			log.Println(err)
		}
		buyers = append(buyers, buyer)
	}
	buyelrsBytes, _ := json.MarshalIndent(buyers, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(buyelrsBytes)

	defer stmt.Close()
	defer db.Close()
}

func FindBuyerById(w http.ResponseWriter, r *http.Request) {
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

	stmt := `SELECT * FROM buyers WHERE id=$1`
	if err != nil {
		log.Println(err)
	}

	row := db.QueryRow(stmt, emp)

	var buyer Buyer
	err = row.Scan(&buyer.ID, &buyer.Name, &buyer.Phone, &buyer.Address)
	if err != nil {
		log.Println(err)
	}

	buyersBytes, _ := json.MarshalIndent(buyer, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(buyersBytes)

	defer db.Close()
}

func EditBuyer(w http.ResponseWriter, r *http.Request) {
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

	var buyer Buyer
	err = json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt := `UPDATE buyers SET name=$1, phone=$2, address=$3 WHERE id=$4`
	if err != nil {
		log.Println(err)
	}

	db.QueryRow(stmt, buyer.Name, buyer.Phone, buyer.Address, emp)

	w.Write([]byte("Update"))

	defer db.Close()
}

func AddBuyer(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	if r.Method != "POST" {
		w.Write([]byte("choose correct method"))
		return
	}

	var buyer Buyer
	err = json.NewDecoder(r.Body).Decode(&buyer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO buyers (name, phone, address) VALUES ($1, $2, $3)")
	stmt.Exec(buyer.Name, buyer.Phone, buyer.Address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("bad requset:", err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
	defer stmt.Close()
}

func RemoveBuyer(w http.ResponseWriter, r *http.Request) {
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

	stmt := `DELETE FROM buyers WHERE id=$1`
	if err != nil {
		log.Println(err)
	}

	db.QueryRow(stmt, emp)

	w.Write([]byte("deleted"))

	defer db.Close()
}
