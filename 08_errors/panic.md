# Panic

So far, you have probably seen your program `panic()` if it can't execute the code because of a syntax error at run time.

These unpredictable errors occur without explicit instructions. Examples include:

- Running out of memory.
- A bad url endpoint
- Trying to divide by 0

A panic will cause the program to stop execution, unless you've included some fail safes within your code.

Let's talk about what happens when a program Panics.

## Panic & Defer

Go has a few built in helper functions that help guide a go program to behave a certain way:

- `defer`: executes a line of code last within a function
- `panic`: called during a run time error, halts execution of the program
- `recover`: tells go what to do after a panic

### Defer

Waits until everything else within a function is complete before firing.

Useful for things like closing a file after doing something with the OS.

Multiple defers are executed in a `LIFO` order.

```go
func doThings() {
  defer fmt.Println("First Line but do this last!")
  defer fmt.Println("Do this second to last!")
  fmt.Println("Things And Stuff should happen first")
}

func main() {
  doThings()
}
```

An example use case would be when opening and closing a file.

```go
f, _ := os.Open(filename)
defer f.Close()

// a bunch of other code you want to run once you've opened the file
```

This is helpful because it keeps two functions that are closely related physically close to each other for readability.

Note: Deferred functions run last, but they are also run _even if_ a runtime panic occurs.

### Panic/Recover

`Panic` will be called during a run time error and fatally kill the execution of a
program.

`Recover` will tell Go what to do when that happens, returning what was passed to `panic`.

Note that in order for recover to do its job, it _must_ be paired with `defer`, which will fire even after a `panic`, otherwise `panic` completely shuts down the execution of a program.

```go
func doThings() {
  for i := 0; i < 5; i++  {
    fmt.Println(i)
    if i == 2 {
      panic("PANIC!")
    }
  }
}

func main() {
  doThings()
}
```

With the above code, once we hit the `panic()` function, our program will stop executing.

Adding a `recover()` cleanup function will tell our program what to do when this happens.

```go
func handlePanic() string {
	return "HANDLING THE PANIC"
}

func recoverFromPanic() {
  // recover() will only return a value if there has been a panic
  if r:= recover(); r != nil {
    fmt.Println("We panicked but everything is fine.")
	  fmt.Println("Panic instructions received:", r)
  }
}

func doThings() {
  defer recoverFromPanic()
  for i := 0; i < 5; i++  {
    fmt.Println(i)
    if i == 2 {
      panic(handlePanic())
    }
  }
}

func main() {
  doThings()
}
```

> _TRY IT_
> Comment out the defer function
> What happens?
