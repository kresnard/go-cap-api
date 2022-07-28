package domain

import (
	"capi/errs"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:inisandi@localhost/banking?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	row := d.client.QueryRow(query, customerID)

	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {

		if err == sql.ErrNoRows {
			log.Println("error customer data not found", err.Error())
			return nil, errs.NewNotFoundError("error customer data not found")
		} else {
			log.Println("error scanning customer data", err.Error())
			return nil, errs.NewUnexpectedError("Enexpected database error")
		}
	}

	return &c, nil
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, *errs.AppErr) {

	query := "select * from customers"
	rows, err := d.client.Query(query)
	if err != nil {
		log.Println("error query data to customer table", err.Error())
		return nil, errs.NewUnexpectedError("error query data to customer table")
	}

	var customers []Customer
	for rows.Next() {

		var c Customer
		rows.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("error scanning customer data", err.Error())
			return nil, errs.NewUnexpectedError("error scanning customer data")
		}
		customers = append(customers, c)
	}
	return customers, nil
}
