package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	"github.com/rhnauf/gokomodo-be/external/db"
	"github.com/rhnauf/gokomodo-be/internal/api"
	"github.com/rhnauf/gokomodo-be/internal/helper"
)

func run() error {
	log.Println("starting up application")

	// load env
	if err := godotenv.Load(); err != nil {
		return errors.New("error loading env files")
	}

	// connect to db
	dbClient, err := db.NewDatabase()
	if err != nil {
		return errors.New(err.Error())
	}

	// seed table buyers & sellers
	//err = seedBuyersTable(dbClient.Client)
	//if err != nil {
	//	return err
	//}

	//err = seedSellersTable(dbClient.Client)
	//if err != nil {
	//	return err
	//}

	httpHandler := api.NewHandler(dbClient)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func seedBuyersTable(db *sqlx.DB) error {
	qry := `
		INSERT INTO buyers (email, name, password, address_send)
		VALUES
		($1, $2, $3, $4)
	`

	stmt, err := db.Prepare(qry)
	if err != nil {
		return err
	}

	for i := 1; i <= 3; i++ {
		email := fmt.Sprintf("buyer_%d@gmail.com", i)
		name := fmt.Sprintf("buyer_%d", i)
		password := fmt.Sprintf("buyer%d", i)
		addressSend := fmt.Sprintf("Jakarta_%d", i)

		hashedPassword, err := helper.HashPassword(password)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(email, name, hashedPassword, addressSend)
		if err != nil {
			return err
		}
	}

	log.Println("buyers table seeded successfully")
	return nil
}

func seedSellersTable(db *sqlx.DB) error {
	qry := `
		INSERT INTO sellers (email, name, password, address_pickup)
		VALUES
		($1, $2, $3, $4)
	`

	stmt, err := db.Prepare(qry)
	if err != nil {
		return err
	}

	for i := 1; i <= 3; i++ {
		email := fmt.Sprintf("seller_%d@gmail.com", i)
		name := fmt.Sprintf("seller_%d", i)
		password := fmt.Sprintf("seller%d", i)
		addressSend := fmt.Sprintf("Bandung_%d", i)

		hashedPassword, err := helper.HashPassword(password)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(email, name, hashedPassword, addressSend)
		if err != nil {
			return err
		}
	}

	log.Println("sellers table seeded successfully")
	return nil
}
