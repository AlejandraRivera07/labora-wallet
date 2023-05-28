package main

import (
	"conectar_db_api/config"
	"conectar_db_api/controllers"
	"conectar_db_api/services"
	"log"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	services.UpDb() //Esta llamada será para levantar la DB, la veremos mas adelante

	//Instancia del Router con Mux
	router := mux.NewRouter()

	//Endpoints
	router.HandleFunc("/items", controllers.GetWallets).Methods("GET")
	router.HandleFunc("/items/{id}", controllers.GetWalletById).Methods("GET")
	router.HandleFunc("/items/{id}", controllers.UpdateWalletByID).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.DeleteWallet).Methods("DELETE")
	router.HandleFunc("/items", controllers.CreateWallet).Methods("POST")
	//services.Db.PingOrDie() //Checkeador de la DB, mas adelante veremos esto
	//handler := cors.Default().Handler(router)
	corsAllowed := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST", "DELETE"},
	})
	handler := corsAllowed.Handler(router)

	//levantamos el servidor
	port := ":9000"
	if err := config.StartServer(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
