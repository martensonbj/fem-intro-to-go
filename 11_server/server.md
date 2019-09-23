# Routes

When you write a Go webapp, you need to pick web "middleware" and a multiplexer (mux) to handle routing. If you're used to things like Ruby on Rails or Django, they come with their own multiplexer built in. Also known as a router or URL router.

To set up routes in Go, we will start by working with the `net/http` package that ships with Go. This package provides the needed interfaces that allow us to read and write to the browser.

As a sidenote, there are many different multiplex packages to deal with http requests in Go. Their job is to match the incoming URL of each request, against a list of registered patterns that you have defined.

Think: Routes in react. If your react code sees the `/` route, it knows to load a certain set of components.

ServeMux is the HTTP request multiplexer that ships with the `http` package found in go source code. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

That being said, it's also touted as a router that contains a certain amount of magic in its routing decisions depending on what the exact pattern it's matching contains.

An example of another common library used is `github.com/gorilla/mux`.

## Defining Routes

Routes are coralled using the `http` package, and the method `HandlFunc`.

```go
func main() {
  http.HandleFunc("/", home)
}
```

Here in our `main` function, `HandleFunc` takes two arguments - the filepath to watch for, and which handler method to call when that url is intercepted.

Before we write out the `home` route handler, let's finish spinning up our web server by adding a few extra lines in `main()`.

```go
func main() {
  http.HandleFunc("/", home)
  fmt.Println("Server is running on port 8080")
  err := http.ListenAndServe("localhost:8080", nil)
  if err != nil {
    return fmt.Errorf("Error firing up server: %v", err)
  }
}
```

## Route Handlers

Route handlers are functions that tell your program what to do when an http request matches a particular pattern. 

```go
func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello World</h1>")
}
```

> Try It
> Add another route to an about page that renders different HTML.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}
```

## Structs & JSON

As front end developers, we are used to working with JSON objecst when we make HTTP requests. 

When working with a struct in Go, we can tell our program what we want our JSON too look like, and then encode/decode it (read: stringify/parse) back and forth into Go.

Let's revisit our User struct.

In order to turn this struct into a JSON object, we can use the `enconding/json` package that ships with Go.

```go
import "encoding/json"

type User struct {
  ID int
  FirstName string
  LastName string
  Email string
}

func main() {
  u := User{
    ID: 1,
    FirstName: "Marilyn",
    LastName: "Monroe",
    Email: "marilyn.monroe@gmail.com",
  }

  data, _ := json.Marshal(u)
  fmt.Println(string(data))
}
```

When you print this out, the structure is ALMOST there, but the formatting isn't conventional JSON.

To modify what we want those keys to look like , we can add an additional field to our struct called a `field tag`, to tell our Struct "Hey, if you're being formatted into json, use this guy instead".

```go
type User struct {
  ID int `json:"id"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  Email string `json:"emailAddress"`
}
```
