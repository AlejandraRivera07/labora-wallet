package services

import (
	"conectar_db_api/models"
	"database/sql"
	"fmt"
	"log"
	"sync"
)

var mu sync.Mutex

// GetItems obtiene todos las wallet de la tabla 'wallet' de la base de datos.
// Retorna una lista de struct 'models.Item' y un error en caso de que haya ocurrido alguno.
func GetWallet() ([]models.Wallet, error) {
	wallets := make([]models.Wallet, 0)
	rows, err := Db.Query("SELECT * FROM wallet")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	// Itera sobre cada fila en 'rows' y crea una instancia de 'models.Item' con los valores de cada columna.
	for rows.Next() {
		var wallet models.Wallet
		err := rows.Scan(&wallet.ID, &wallet.CustomerId, &wallet.CreateDate, &wallet.CountryId)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		wallets = append(wallets, wallet)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return wallets, nil
}

func GetWalletById(id int) (*models.Wallet, error) {

	query := `
    SELECT 	*
    FROM wallet 
	WHERE id = $1;
    `
	var wallet models.Wallet

	err := Db.QueryRow(query, id).Scan(&wallet.ID, &wallet.CustomerId, &wallet.CreateDate, &wallet.CountryId)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontró ningún objeto
			return nil, nil
		}
		return nil, err
	}
	return &wallet, nil
}

func CreateWallet(wallet models.Wallet) (models.Wallet, error) {
	_, err := Db.Exec("INSERT INTO items (id, customer_id, create_date, country_id)",
		wallet.CustomerId, wallet.CreateDate, wallet.CountryId)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return models.Wallet{}, err
	}

	return wallet, nil
}

func UpadateWalletById(id int, wallet models.Wallet) (models.Wallet, error) {

	query := `UPDATE wallet SET customer_id=$1, create_date=$2, country_id=$3 WHERE id=$1`

	// execute the sql statement
	_, err := Db.Exec(query, wallet.CustomerId, wallet.CreateDate, wallet.CountryId)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return models.Wallet{}, err
	}
	return wallet, nil

}

func DeleteWalletById(id int, wallet models.Wallet) (models.Wallet, error) {

	query := `DELETE FROM wallet WHERE id=$1`

	// execute the sql statement
	_, err := Db.Exec(query, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return models.Wallet{}, err
	}
	return wallet, nil
}
