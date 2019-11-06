# Control Structures

## If Statements

If statements look similar to JavaScript, however in Go you exclude the
parentheses around the conditional statement.

An example if statement:

```go
  if someVar > 10 {
    // do some stuff
  }
```

Similarly, you have `else if` and `else` blocks.

```go
  if i > 100 {
    fmt.Println("Greater than 100")
  } else if i == 100 {
    fmt.Println("Equals 100")
  } else {
    fmt.Println("Less than 100")
  }
```

### An Intro to Error Handling

We will discuss error handling in depth later in this course, but lets dive in
a bit now.

First of all, a noticeable aspect of Go is the fact that as the developer you
need to explicitly catch errors, and do something about them if they exist. One
such approach might look like this:

```go
err := someFunction()
// => If this function returns a value, it will be an  error of type Error

if err != nil {
  fmt.Println(err.Error())
}
```

This is a very common situation, and Go let's us combine these lines into the
same if block:

```go
if err := someFunction(); err != nil {
  fmt.Println(err.Error())
}
```
### If Blocks & Scope

*PRO TIP* The `if` keyword kicks off it's own variable scope.

This means that the variable `err` is scoped to this if block, which is helpful since we
will be using this type of syntax all over the place.

## Switch

An example switch statement:

```go
  var city string

  switch city {
  case "Des Moines":
    fmt.Println("You live in Iowa")
  case "Minneapolis,", "St Paul":
    fmt.Println("You live in Minnesota")
  case "Madison":
    fmt.Println("You live in Wisconsin")
  default:
    fmt.Println("You're not from around here.")
  }
```

Note that you can use multiple arguments in your case statement separated by a
comma.

A second version can contain no `switch expression` (represented by `city` after
the switch keyword in the above statement)

```go
var i int

switch {
  case i > 10: fmt.Println("Greater than 10")
  case i < 10: fmt.Println("Less than 10")
  default: fmt.Println("Is 10")
}
```

Note that in Go, if a case statement is executed, Go will not automatically
continue running the switch statement. To ask for this behavior, you can add a
`fallthrough` keyword.

```go
var i int = 9

switch {
  case i != 10:
    fmt.Println("Does not equal 10")
    fallthrough
  case i < 10: fmt.Println("Less than 10")
  case i > 10: fmt.Println("Greater than 10")
  default: fmt.Println("Is 10")
}
```

## For Loops

The only loop in Go is the `for` loop. It looks nearly identical to our friendly for loop in JavaScript, but depending on how we write it it can behave in various ways.

Like variables, we'll start with the most verbose version and work our way to idiomatic go.

```go
func main() {
  i := 1
  for i <= 100 {
    fmt.Println(i)
    i += 1
  }
}
```

Like in JS, we instantiate our counter `i` at 1, using shorthand variable assignment. Note that we do not need parentheses around the conditional.

For every iteration, as long as our counter `i` is below 100, we print its value

Increment the counter by one.

We can also use a shorter way to write this that looks more like the version we all know and
love in JS:

```go
  func main() {
    for i := 1; i <= 100; i++ {
      fmt.Println(i)
    }
  }
```

If we omit the first and third pieces of the opening statement (variable
declaration and increment), the for loop will behave like a `while` loop:

```go
  func main() {
    for i < 100 {
      // do some stuff
    }
  }
```

### Iterating Over A List

For loops also let us iterate over a collection of elements, similar to how
`forEach` might feel in JS (although we are not limited to an array).

To do this we use the `range` keyword.

```go
  for index, element := range someCollection {
    // Do some stuff
  }
```

## Exercise 3a: Control Structures
`exercise_3a.md`
