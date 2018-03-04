package main_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"."
)

var a main.App

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func TestMain(m *testing.M) {
	a = main.App{}

	os.Setenv("TEST_DB_USERNAME", "gopher")
	os.Setenv("TEST_DB_PASSWORD", "gopher")
	os.Setenv("TEST_DB_NAME", "go_test")

	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"))

	fmt.Println("TEST_DB_USERNAME:", os.Getenv("TEST_DB_USERNAME"))
	fmt.Println("TEST_DB_PASSWORD:", os.Getenv("TEST_DB_PASSWORD"))
	fmt.Println("TEST_DB_NAME:", os.Getenv("TEST_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
