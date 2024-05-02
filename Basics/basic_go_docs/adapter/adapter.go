package adapter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"user_id" xml:"userID"`
	ID        int    `json:"id" xml:"id"`
	Title     string `json:"title" xml:"title"`
	Completed bool   `json:"completed" xml:"completed"`
}

// ADAPTER
type DataInterface interface {
	GetData() (*Todo, error)
}

type RemoteService struct {
	Remote DataInterface
}

func (rs *RemoteService) CallRemoteService() (*Todo, error) {
	return rs.Remote.GetData()
}

type JSONBackend struct{}

func (jb *JSONBackend) GetData() (*Todo, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo Todo
	err = json.Unmarshal(body, &todo)

	if err != nil {
		return nil, err

	}
	return &todo, nil
}

type XMLBackend struct{}

func (xb *XMLBackend) GetData() (*Todo, error) {
	// Faking
	xmlFile := `
	<?xml version="1.0" encoding="UTF-8"?>
	<root>
		<userID>1</userID>
		<id>1</id>
		<title>delectus aut autem</title>
		<completed>false</completed>
	</root>
	`
	var todo Todo
	_ = xml.Unmarshal([]byte(xmlFile), &todo)
	return &todo, nil
}

//

func TestAdapter() {
	// Wihout adapter
	todo := getRemoteData()
	fmt.Println("todo without adapter pattern:", todo)
	// with adapter JSON
	jsonBackend := &JSONBackend{}
	jsonAdapter := &RemoteService{Remote: jsonBackend}
	todoFromJSON, _ := jsonAdapter.CallRemoteService()
	fmt.Println("from  json adapter :", todoFromJSON)
	// with adapter XML
	xmlBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: xmlBackend}
	todoFromXML, _ := xmlAdapter.CallRemoteService()
	fmt.Println("from  xml adapter :", todoFromXML)

}

func getRemoteData() *Todo {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var todo Todo
	err = json.Unmarshal(body, &todo)

	if err != nil {
		log.Fatal(err)
	}
	return &todo
}
