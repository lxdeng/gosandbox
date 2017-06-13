package main

import (
	"database/sql"
	"errors"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
	return err
}

func (p *product) createProduct(db *sql.DB) error {
	/*
		SQLite does not suport the DML Returning statement
		err := db.QueryRow(
			"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
			p.Name, p.Price).Scan(&p.ID)
	*/

	stmt, err := db.Prepare("INSERT INTO products(name, price) VALUES(?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(p.Name, p.Price)
	if err != nil {
		return err
	}

	if id, err := res.LastInsertId(); err != nil {
		return err
	} else {
		p.ID = int(id)
	}

	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}