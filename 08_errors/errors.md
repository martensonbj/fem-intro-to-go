# Error Type

The first type of error we will talk about is the built in `error` type. 

When this occurs, most of the time you will have a game plan for what to do instead and your program can continue to run.

The built in `error` type is an `interface` type (which will be discussed shortly), with a method called `Error()` that returns the string of the error message itself.

That interface looks like this:

```go
type error interface {
    Error() string
}
```

```go
func isGreaterThanTen(num int) error {
  if num < 10 {
    return fmt.Errorf("%v is NOT GREATER THAN TEN", num)
  }
  return nil
}

func main() {
  err := isGreaterThanTen(9)
  if err != nil {
    log.Println(err)
  } else {
    fmt.Println("Carry On.")
  }
}
```

Here, the `fmt` package will return an error. (Run `go doc fmt.Errorf` to check it out).

IMPORTANT NOTE: `nil` is an acceptable return value for an `error` type as well.

NOTE: Log: This is a simple logging package built in to Go. You can format the output similarly to `fmt` which prints it to the console. Log will print the date and time of each logged message. Again, check out `go doc log` for more details.

_TRY IT_
Refactor the main function to consolidate the variable `err` declaration and the if check.

Hint: Check out the section on `if` statements.

## Exiting The Program

If your program encounters an error that should halt execution, you can use `log.Fatal` from the `log` package which will log the error, and then exit the program.

```go
func isGreaterThanTen(num int) error {
  if num < 10 {
    return fmt.Errorf("%d is NOT GREATER THAN TEN", num)
  }
  return nil
}

func main() {
  err := isGreaterThanTen(9)
  if err != nil {
    log.Fatalln(err)
  } else {
    fmt.Println("Carry On.")
  }
}
```

## Error Handling in Go 1.13

We will cover the newest approach to error handling released in Go 1.13 towards the end of the course 

