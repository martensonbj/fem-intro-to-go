## Methods

Similar to structs and variables, functions can also be bound to types.

Let's revisit the `describeUser` function , and User struct, from a few examples ago:

```go
type User struct {
  ID int
  FirstName string
  LastName string
  Email string
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.firstName, u.lastName, u.email, u.ID)
	return desc
}

func main() {
  user := User{ID: 1, firstName: "Marilyn", lastName: "Monroe", email: "marilyn.monroe@gmail.com"}
  desc := describeUser(user)
}
```

In the function signature above, we name our function `decribeUser()`, and pass in an explicit argument of type `User`, which must return a `string`.

To refactor this into a method, instead of passing a `User` struct as an argument, we insert it as a `receiver` between the `func` keyword and the name of our slightly renamed `describe()` function.

This tells Go that when the `describe()` function receives (is called on) a struct with the shape `User`, it should execute the code within the curly braces.

We can then call the method a little differently:

```go
func (u *User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s, ID: %d", u.firstName, u.lastName, u.email, u.ID)
	return desc
}

func main() {
  u := User{ID: 1, firstName: "Marilyn", lastName: "Monroe", email: "marilyn.monroe@gmail.com"}
  desc := u.describe()
}
```

Note that methods DO NOT require the `&` operator to permanently modify the struct being acted on- Go automatically passes a pointer to a method call.

## Advantages

- Encapsulation
- Polymorphism - two different types can have a method of the same name (as seen
  above)
- Functions that control/modify state

## Exercise 9
`09_methods/exercise_9a.md`
