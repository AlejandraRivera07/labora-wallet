package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"conectar_db_api/models"
	"conectar_db_api/services"

	"github.com/gorilla/mux"
)

type WalletHandler struct {
	ItemService services.DbConnection
}

func GetWallets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := services.GetWallet()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal Server Error")
		return
	}
	json.NewEncoder(w).Encode(items)

}

func GetWalletById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Fatal(err)
	}

	item, err := services.GetWalletById(id)

	if item == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Objeto con id %d no encontrado", id)))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

func CreateWallet(w http.ResponseWriter, r *http.Request) {
	// create an empty wallet of type models.User
	var walletCreate models.Wallet

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&walletCreate)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	// call insert wallet function
	walletToCreate, err := services.CreateWallet(walletCreate)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("wallet created: %+v", walletToCreate)
	jsonData, err := json.Marshal(walletToCreate)
	if err != nil {
		log.Printf("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))

}

func UpdateWalletByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty wallet of type models.Wallet
	var walletUpdate models.Wallet

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&walletUpdate)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update function
	walletUpdated, err := services.UpadateWalletById(id, walletUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("wallet updated: %+v", walletUpdated)
	jsonData, err := json.Marshal(walletUpdated)
	if err != nil {
		log.Printf("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}

func DeleteWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get the walletid
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	var idWalletDelet models.Wallet
	// call the delete function,
	deletedWallet, err := services.DeleteWalletById(id, idWalletDelet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("wallet updated: %+v", deletedWallet)
	jsonData, err := json.Marshal(deletedWallet)
	if err != nil {
		log.Printf("Error %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonData))
}
