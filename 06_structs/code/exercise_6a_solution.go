package main

// import "fmt"

// // User is a user type
// type User struct {
// 	ID                         int
// 	FirstName, LastName, Email string
// }

// // Group represents a set of users
// type Group struct {
// 	role           string
// 	users          []User
// 	newestUser     User
// 	spaceAvailable bool
// }

// func describeUser(u User) string {
// 	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
// 	return desc
// }

// // func describeGroup
// // => "This user group has 19 users. The newest user is Joe Smith. Accepting New Users: true"

// func describeGroup(g Group) string {

// 	if len(g.users) > 2 {
// 		g.spaceAvailable = false
// 	}

// 	desc := fmt.Sprintf("This user group has %d. The newest user is %s %s. Accepting New Users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
// 	return desc
// }

// func main() {
// 	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}

// 	u2 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

// 	u3 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

// 	g := Group{
// 		role:           "admin",
// 		users:          []User{u, u2, u3},
// 		newestUser:     u2,
// 		spaceAvailable: true,
// 	}

// 	fmt.Println(describeUser(u))
// 	fmt.Println(describeGroup(g))
// 	fmt.Println(g)
// }
