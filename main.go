package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Customers struct {
	CustomerID   int    `json:"customerid"`
	CustomerName string `json:"customername"`
	CustomerAge  int    `json:"customerage"`
}

func main() {

	fmt.Println("Web Services and Cloud Services Using Golang and MySQL")

	db, err := sql.Open("mysql", "adminweb:testadmin@tcp(webcloudservices.c3ayqo6cawed.us-east-1.rds.amazonaws.com:3306)/webcloudservices")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Successfully Connected to MySQL")

	insert, err := db.Query("INSERT INTO customers VALUES(1,'Minions', 20)")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Successfully inserted into customers tables")
	insert2, err := db.Query("INSERT INTO customers VALUES(2,'Gru', 15)")
	if err != nil {
		panic(err.Error())
	}
	defer insert2.Close()
	fmt.Println("Successfully inserted into customers tables")
	// ini jika mau delete
	// delete, err := db.Query("DELETE FROM customers WHERE CustomerName LIKE 'Minions'")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer delete.Close()
	// fmt.Println("Successfully deleted from customers tables")
	http.HandleFunc("/viewdata", func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT CustomerID, CustomerName, CustomerAge FROM customers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var tempQuery []Customers

		for rows.Next() {
			var tempCust Customers
			if err := rows.Scan(&tempCust.CustomerID, &tempCust.CustomerName, &tempCust.CustomerAge); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tempQuery = append(tempQuery, tempCust)
		}
		if err := rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tempQuery)

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Web and Cloud Services")
	})
	http.ListenAndServe(":8000", nil)

}
