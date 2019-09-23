package main

import "fmt"

// User is a user type
type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
}

// // User is a user type
// type User struct {
// 	ID                         int
// 	FirstName, LastName, Email string
// }

func main() {
	var u User
	// u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

	fmt.Println(u)
}
