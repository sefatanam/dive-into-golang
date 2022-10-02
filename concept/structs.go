package main

import "fmt"
type User struct {
	FirstName string`user:"first_name"`
	LastName string`user:"last_name"`
}

func (u *User) String() string{
	return fmt.Sprintf("First Name: %s", u.FirstName)
}

// func main(){
// 	u:= &User{
// 		FirstName: "John",
// 		LastName: "Doe",
// 	}
// 	fmt.Println(u)
// }