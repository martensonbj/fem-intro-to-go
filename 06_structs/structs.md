# Defining a struct

A `struct` is a data type that contains a collection of fields that act as variables, have defined types and is reusable across your program.

Think an ES6 Class in JS, or an object of values without keys.

```go
type User struct {
  ID int
  FirstName string
  LastName string
  Email string
}
```

Another way to write this is:

```go
type User struct {
  ID int
  FirstName, LastName, Email string
}
```

We can now treat the `User` struct as a type of variable and instantiate it with or without default values.

```go
var u User
fmt.Pringln(u)
```

What do you notice about the initial values of the variable `u`?

Let's add values.

```go
var u User
u.ID = 1
u.FirstName = "Marilyn"
u.LastName = "Monroe"
u.Email = "marilyn.monroe@gmail.com"
```

Or, more commonly, in shorthand

```go
u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
```


# Reading a Struct

Then, in order to access the fields, we use dot notation similar to working with
Objects in JS.

```go
u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
fmt.Println(u.FirstName) // => "Marilyn"
```

To pass a struct to a function, we simply add the struct name within the
function signature as the type of argument we are passing in.

```go
func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.FirstName, u.LastName, u.Email, u.ID)
	return desc
}
```

In `main()`, add a few lines to call this function.

```go
func main() {
  u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
  userDescription := describeUser(u)
  fmt.Println(userDescription)
}
```

Let's create another struct for a group of users (think `admin` vs `guest`).

```go
type Group struct {
	role            string
	users           []User
  newestUser      User
  spaceAvailable  bool
}
```

Next, create a function that describes the team of users:

```go
func describeGroup(g Group) string {
	desc := fmt.Sprintf("The %s user group has %d users. Newest user:  %s, Accepting New Users:  %t", g.role, len(g.users), g.newestUser.FirstName, g.spaceAvailable)
	return desc
}
```

In main we can set up instances of our structs and call each function:

```go
func main() {
  u1 := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
  u2 := User{ID: 2, FirstName: "Humphrey", LastName: "Bogart", Email: "humphrey.bogart@gmail.com"}

  g := Group{role: "admin", users: []User{u1, u2}, newestUser: u2, spaceAvailable: true}

  userDescription := describeUser(u1)
  groupDescription := describeGroup(g)

  fmt.Println(userDescription)
  fmt.Println(groupDescription)
}
```

## Accessibility

Note that like variables and functions, only field names that have a capital
letter are visible outside of the package where the struct is defined.

```go
type User struct {
  ID  int
  Email  string
  FirstName  string
  LastName   string
}
```

## EXERCISE 6A
`06_structs/exercise_6a`


# Mutating Instances Of A Struct

In Go, unless otherwise indicated, we are passing around a _COPY_ of any variable we reference to existing functions. In the case above, we pass a copy of the variable `g`, and then are asking our program to print the original value of `g`. The original value is still safely untouched in memory.

In order to permanently modify this variable, we need to use a special kind of variables called a `pointer`.