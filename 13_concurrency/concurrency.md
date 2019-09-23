# Concurrency

## What is Concurrency in Go?

## Sync

Goroutine: From [the docs](https://tour.golang.org/concurrency/1), a goroutine is a lightweight threat managed by the Go runtime.

A goroutine is indicated by adding the keyword `go` before the name of a function.

This will tell Go to fire up a new goroutine running that function on a separate thread.

```go
  import (
    "time"
    "fmt"
  )

  func say(s string) {
    for i := 0; i < 3; i++ {
      fmt.Println(s)
      time.Sleep(time.Millisecond*100)
    }
  }

  func main() {
    go say("Hello")
    say("There")
  }
```

If all of the code within a function is a go routine:

```go
  func main() {
    go say("Hello")
    go say("There")
  }
```

Everything will be non blocking, and nothing will finish execution. In order to fix this, we need to synchronize our goroutines, using the package [sync](https://golang.org/pkg/sync/).

To start, create a variable that sets up a `WaitGroup` - meaning a set of go routines you want to execute to completion before moving forward in your program.

```go
package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func handlePanic() {
  if r := recover(); r != nil {
    fmt.Println("We panicked! But its fine. Message received: ", r)
  }
}

func printStuff(s string) {
  // Decrement the wait group counter
  // Use defer so that if the function panics we aren't waiting forever
  // Also figure out what to do if the program panics
	defer wg.Done()
  defer handlePanic()

  for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
  // Increment the wait group counter
	wg.Add(1)
  // Launch a goroutine
	go printStuff()
  // Increment the wait group counter
	wg.Add(1)

	go printStuff()
  // Wait for the waitgroup counter to be 0 before continuing
	wg.Wait()
}
```

Here, we're defining a variable that represents our WaitGroup. The `Add()`
function takes an integer and adds it to a behind the scenes counter. If the
counter becomes zero, all goroutines are released and the program will continue.

## Channels

Send and receive values between your goroutines.

Create a channel but using the `make()` function, and the `chan` keyword,
including the type declaration.

```go
myChannel := make(chan int)
```

Then, fire off a couple goroutines form which you are expecting values.

```go
package main

import "fmt"

func mathThings(c chan int, someValue int) {
	// From right to left, calculate the value, send it to the channel
	c <- someValue * 5
}

func main() {
// Instantiate a new channel
	resultChan := make(chan int)

	// Kick off a couple go routines
	// Notice that we don't need the WaitGroup here, because channels are blocking.
	// Our channel is expecting a value back so it will block until it finishes execution
	go mathThings(resultChan, 5)
	go mathThings(resultChan, 3)

	// Store whatever we get back from our channel, in a var
	num1 := <-resultChan
	num2 := <-resultChan
  // or num1, num2 := <-resultChan, <-resultChan
	fmt.Println(num1, num2)
}
```

We can shorten this up a bit by iterating over our channel, adding back in
WaitGroups, and including a buffer value when instantiating our channel:

```go
package main

import
  "fmt"
  )

var wg sync.WaitGroup

func mathThings(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	// second argument is the buffer to tell our channel how many times we plan on using it
	resultChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go mathThings(resultChan, i)
	}

	// Wait for all goroutines to complete before closing the channel
	wg.Wait()

	// Close wont wait for the goroutines to be done without additional synchronization
	close(resultChan)


	for item := range resultChan {
		fmt.Println(item)
	}
}
```
