# Interfaces

If you think of structs as a set of properties that define a type, interfaces can be thought of as a set of methods that define a type.

In other words, an interface contains a list of function signatures, describing the behavior of other types.

```go
type Describer interface {
  describe() string
}
```

Once again we start with the `type` keyword, followed by the name of the
interface, and then the actual keyword `interface`.

Note that the convention is to name an interface with the "job" it is doing - often ending in `er`.

Within the curly braces we list a set of methods that are associated with our interface type and the type those methods should return.

Our interface `Describer` includes any type that has a method named `describe()`, which in our program includes `User` and `Group`. Any type that defines this method is said to `satisfy the Describer interface`.


```go
func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

	describeUser := u1.describe()
	describeGroup := g.describe()

	fmt.Println(describeUser)
	fmt.Println(describeGroup)
}

```

You can also use the empty interface type to indicate that Go should accept
anything.

```go
interface{}
```

```go
func DoSomething(v interface{}) {
  // This function will take anything as a parameter.
}
```

Once again, our entire program should now look like this:

```go
package main

import "fmt"

// Describer prints out a entity description
type Describer interface {
	describe() string
}

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

func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}

func (g *Group) describe() string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}

// DoTheDescribing can be implemented on any type that has a describe method 
func DoTheDescribing(d Describer) string {
	return d.describe()
}

func main() {
	u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 1, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}
	g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

	describeUser := u1.describe()
	describeGroup := g.describe()

	userDescribeWithIface := DoTheDescribing(&u1)
	groupDescribeWithIface := DoTheDescribing(&g)

	fmt.Println(describeUser)
	fmt.Println(describeGroup)

	fmt.Println(userDescribeWithIface)
	fmt.Println(groupDescribeWithIface)
}
```
