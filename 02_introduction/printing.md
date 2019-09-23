# Printing

We've seen `fmt.Println` used to effectively log something the console.

Note that this function can take multiple arguments, and will insert a space between
arguments (which differs from having to manually insert them in JS 🙄).

`fmt.Println()`

```go
package main

import "fmt"

func main() {
	fmt.Println("All the best words start with B")
	fmt.Println("Bears, Beets, Battlestar Gallactica")
}
```

`fmt.Print()`: Does not insert a new line

```go
package main

import "fmt"

func main() {
	fmt.Print("All the best words start with B")
	fmt.Print("Bears, Beets, Battlestar Gallactica")
}
```

`fmt.Printf()`: Uses string formatting to interpolate variables
Note:
`%t`: expects a boolean
`%d`: expects a integer
`%s`: expects a string

```go
package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, it is %t that I have %d brother.", "Marilyn", true, 1)
}
```

## Similar Functions and Patterns

You'll notice that many of the functions that log or print will have similar
flags with different underlying purpose.

`fmt.Print()`

- Prints output to the stdout console
- Returns number of bytes and an error, although error is generally not
  worried about

`fmt.Fprint()`

- Prints the output to an external source (file, browser), not on the stdout console
- Returns number of bytes, and any write errors

`fmt.Sprint()`:

- Stores output on a character buffer instead of spitting it out to the console
- Returns the string you want to print

## Exercise 2a: Hello World

`exercise_2a.md`
