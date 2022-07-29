package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:inisandi@localhost/banking?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	// row := d.client.QueryRow(query, customerID)
	// err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)

	var c Customer

	err := d.client.Get(&c, query, customerID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("error customer data not found" + err.Error())
			return nil, errs.NewNotFoundError("error customer data not found")
		} else {
			logger.Error("error scanning customer data" + err.Error())
			return nil, errs.NewUnexpectedError("Enexpected database error")
		}
	}

	return &c, nil
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppErr) {
	query := "select * from customers"

	var customers []Customer

	if status == "" {
		query = "select * from customers"
		err := d.client.Select(&customers, query)
		if err != nil {
			log.Println("error query data to customer table", err.Error())
			return nil, errs.NewUnexpectedError("error query data to customer table")
		}
	} else if status == "active" || status == "inactive" {
		if status == "active" {
			status = "1"
		} else {
			status = "0"
		}
		query = "select * from customers where status = $1"
		err := d.client.Select(&customers, query, status)
		if err != nil {
			log.Println("error query data to customer table", err.Error())
			return nil, errs.NewUnexpectedError("error query data to customer table")
		}
	}

	return customers, nil
}
