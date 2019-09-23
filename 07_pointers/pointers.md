# Pointers

This can be a particularly tricky aspect of Go when coming from a dynamically typed language like JavaScript.

Review: When defining a variable, that variable gets stored in a location in computer memory, which has a memory address.

At the base level, a pointer variable in Go is a variable that holds the _memory location_ of that variable, instead of a copy of its value.

```go
  var name string
  var namePointer *string

  fmt.Println(name)
  fmt.Println(namePointer)
```

You'll notice here that set to their default values, a normal string variable (`name`) is assigned a value of `""`.

A pointer string variable (`namePointer`), however, is assigned it's default value of `<nil>` in anticipation of a future memory location.

To assign a variable it's address in memory, you use the `&`
symbol (think "a - at - for address").

```go
  var name string
  var namePointer *string

  fmt.Println(name)
  fmt.Println(namePointer)
  fmt.Println(&name)
```

_SLIDE_ 

_TRY IT_
Set both `name` and `namePointer` to values.
What happens?
What does the error message mean?
How do you assign an address location of the `name` variable to `namePointer` variable instead?

// SOLUTION

```go
  var name string = "Beyonce"
  var namePointer *string = &name

  fmt.Println(name)
  fmt.Println(namePointer)
```

So now we're looking at the _memory address_ of a variable, rather than the value that is stored there. 

To get that value, we need to dig into that address and pull the value out. 

To `read through` a pointer variable to get the actual value, you use `pointer indirection` to "deference" that variable back to its original value.

```go
  var name string = "Beyonce"

  var namePointer *string = &name
  var nameValue = *namePointer

  fmt.Println(name)
  fmt.Println(namePointer)
  fmt.Println(nameValue)
```

### To Summarize In Other Words:

- Pointer types are indicated with a `*` next to the _TYPE NAME_, indicates that variable will POINT TO a memory location.

- Pointer variable values are visible with a `*` next to the _VARIABLE NAME_.

- To read through a variable to see the pointer address, use a `&` next to the _VARIABLE NAME_.


  - ie: `var nameValue = *namePointer` => "Marilyn"

## Pass By Value

Function parameters represent a COPY of of the value that they reference. 

This means that any change made to that variable within the function is modifying the copy, not the underlying value in memory. 

```go
func changeName(n string) {
    n = strings.ToUpper(n)
}

func main() {
    name := "Elvis"
    changeName(name)
    fmt.Println(name)
}
```

If we did want the `changeName` function to modify the name variable, we need to pass a pointer.

```go
func changeName(n *string) {
	*n = strings.ToUpper(*n)
}

func main() {
	name := "Elvis"
	changeName(&name)
	fmt.Println(name)
}
```
 
 ## Pointers And Structs

As the [go gocs](https://tour.golang.org/moretypes/4) mention, Go gives us a small loophole when working with pointer structs.

Let's take the following struct:

```go
type Coordinates struct {
  X, Y float64
}

var c = Coordinates{x: 120.5, y: 300}
```

To access a field on a struct pointer, we would need syntax that looked like this:

```go
func main() {
  pointerCoords = &c
  (*pointerCoords).X =  200
}
```

The `(*pointerCoords)` bit is a tad cumbersome, so Go allows us to simply reference the pointer without the dereferencing asterisk but knows that behind the scenes we are talking about the pointer struct.


