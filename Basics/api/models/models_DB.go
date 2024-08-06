package models

var todos = []Todo{
	{Id: 1, Name: "Todo 1", UserId: 1, Description: "this is a simple description"},
	{Id: 2, Name: "Todo 2", UserId: 2, Description: "this is a simple description"},
	{Id: 3, Name: "Todo 3", UserId: 1, Description: "this is a simple description"},
	{Id: 4, Name: "Todo 4", UserId: 1, Description: "this is a simple description"},
	{Id: 5, Name: "Todo 5", UserId: 3, Description: "this is a simple description"},
	{Id: 6, Name: "Todo 6", UserId: 3, Description: "this is a simple description"},
	{Id: 7, Name: "Todo 7", UserId: 3, Description: "this is a simple description"},
	{Id: 8, Name: "Todo 8", UserId: 2, Description: "this is a simple description"},
	{Id: 9, Name: "Todo 9", UserId: 4, Description: "this is a simple description"},
	{Id: 10, Name: "Todo 10", UserId: 1, Description: "this is a simple description"},
}

var users = []User{
	{Id: 1, Username: "sefat", Password: "sefat"},
	{Id: 2, Username: "anam", Password: "anam"},
	{Id: 3, Username: "amal", Password: "khan"},
	{Id: 4, Username: "borhan", Password: "khan"},
}

func FakeTodos() *[]Todo {
	return &todos
}

func FakeUsers() *[]User {

	return &users
}
