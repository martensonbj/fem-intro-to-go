# Exercise 3b: (Bonus) Revisiting Hello World

## Goals 

- Practice using the `fmt.Scan()` function
- Practice string formatting with `fmt.Printf()`

## Setup

1. Let's revisit part 2 of the Hello World exercise from section 2:

```go
package main

import "fmt"

func main() {
  fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t", "Brenna", "Denver", 4, true)
}
```

## Directions

Instead of hard coding in these values, let's refactor our code to collect user
input.

2. Using `fmt.Scan()`, ask the user for each of the missing variables, using a new `fmt.Println()` request for each variable:

- Their name
- What city they live in
- How long they have lived there
- If the weather is nice

3. Print the same sentence as part 2 of exercise_2a.

## Hints

_Hint_: Don't forget to include the `&` symbol in the `Scan()` function.

_Hint_: Arguments coming from the terminal will always be strings, might need
to do some converting.

_Hint_: Might need to pull in an additional package to handle string formatting, to make sure a user can type either "yes" or "YES" or "Yes" etc...

