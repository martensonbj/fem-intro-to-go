package main

import "fmt"

// User is a generic user type
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func main() {
	user := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	desc := describeUser(user)
	fmt.Println(desc)
}
