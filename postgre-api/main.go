package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	passowrd = "password"
	dbname   = "office"
)

func getConnection() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, passowrd, dbname)
}

func getEmployees(db *sql.DB) {
	rows, err := db.Query("SELECT id, firstName, lastName , email FROM employees")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var email string

		err := rows.Scan(&id, &firstName, &lastName, &email)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, firstName, lastName, email)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}


func main() {
	connectionString := getConnection()
	db, err := sql.Open("postgres", connectionString)


	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful.")

  getEmployees(db)
}
