package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name    string `json:"name"`
	Age     string `json:"age"`
	Country string `json:"country"`
}

func (person Person) WriteUserInput() error {
	jsonData, err := json.Marshal(person)
	if err != nil {
		return err
	}

	file, err := os.Create(person.Name + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}
	return nil
}

func main() {

	person := Person{}
	fmt.Printf("Enter Your Name: ")
	person.Name = GetUserInput()

	fmt.Printf("Enter Your Age: ")
	person.Age = GetUserInput()

	fmt.Printf("Enter Your Country: ")
	person.Country = GetUserInput()

	err := person.WriteUserInput()

	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/detail", GetUserDetailRequest)
	http.HandleFunc("/submit", PostUserDetailRequest)
	err = http.ListenAndServe("localhost:3000", nil)

	if err != nil {
		log.Println(err)
	}
}

func GetUserInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func GetUserDetailRequest(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "name query param is missing", http.StatusBadRequest)
		return
	}

	jsonData, err := os.ReadFile(name + ".txt")
	if err != nil {
		http.Error(w, "error in reading file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func PostUserDetailRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "request not allowed", http.StatusBadRequest)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "name query param is missing", http.StatusBadRequest)
		return
	}

	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)

	if err != nil {
		http.Error(w, "cannot parse request body", http.StatusBadRequest)
		return
	}
	log.Println(person)

	err = person.WriteUserInput()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
