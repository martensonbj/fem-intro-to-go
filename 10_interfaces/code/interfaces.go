package main

import "fmt"

// Add a Describer interface

// User is a single user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group is a group of Users
type Group struct {
	role           string
	users          []User
	newestUser     User
	spaceAvailable bool
}

// These two structs have different implementations of the `describe()` method.

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

// Create a function that doesn't care what type you pass in as long as the type "satisfies the interface"

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}
	describeUser := u1.describe()
	describeGroup := g.describe()
	fmt.Println(describeUser)
	fmt.Println(describeGroup)
}
