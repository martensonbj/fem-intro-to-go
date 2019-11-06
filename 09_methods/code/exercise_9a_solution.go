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

// func (u *User) describe() string {
// 	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
// 	return desc
// }

// func (g *Group) describe() string {
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

// 	fmt.Println(g.describe())
// 	fmt.Println(u.describe())
// 	fmt.Println(g)
// }
