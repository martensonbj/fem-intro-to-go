# Functions

Functions in Go, like in most languages, map a specified set of inputs to specified outputs.

In Go, the input parameters also require a type definition, with an optional return type definition.

```go
func printAge(age int) int {
  fmt.Println(age)
  return age
}
```
In the above function, we name the function `printAge()`, which expects an input
type of an integer. The value between the closing argument parens and the opening
curly brace is the type this function must return at the end of the day.

## Multiple Returns

A function can also return multiple values, which both need to be specified in
the return type definition:

```go
  func myFunc() (int, int) {
    return 5, 6
  }

  func main() {
    x, y := myFunc()
  }
```

## Named Returns

```go
  func myFunc() (ageOfBob int, ageOfSally int) {
    ageOfBob = 21
    ageOfSally = 25
    // return ageOfBob, ageOfSally
    return
  }

  func main() {
    x, y := myFunc()
  }
```

Note that when calling this function, in order to access both return values you
also need to define two variables.

## Exercise 4a: Functions
`exercise_4a_functions.md`

## Variadic Function

Similar to the spread/rest operator in JavaScript, you can pass an unspecified number
of arguments to a function, all of which are stored as a `slice` of the same type of element, which we
will talk about shortly.

```go
  func doThings(args ...int) int {
    total := 0
    for _, num := range args {
      total += num
    }
    return total
  }

  func main() {
    fmt.Println(doThings(1, 2, 3))
  }
```

Note that the variadic parameter must be the last on in the argument list.

## Exercise 4b: Refactor Functions
`exercise_4b_functions.md`

### Scope

Similar to JS, variables defined within a function can only be accessed within that function.

To use a variable in other functions within the program, you must define it externally.

```go
func otherFunc() {
	fmt.Println(user) // pilot will be undefined
}

func main() {
  var user string = "Marilyn Monroe"
	fmt.Println(user)
}
```

Note: Functions do not have access to variables that are not global, or defined within
themselves.

### Block Scope

Note that curly braces also indicate their own scope. Variables defined within a
set of curly braces are not accessible outside of those curly braces.

This is similar to if blocks (which we saw earlier).

## More On Functions

### Functions As Types

Functions can also be stored as variables, with their function signature as
their type.

```go
func isOlderThanTen(age int) bool {
  return age > 10
}

func isNotZero(num int) bool {
  return num != 0
}

var myFuncVariable func(int) bool

func main() {
  myFuncVariable = isOlderThanTen
  fmt.Println(myFuncVariable(9)) // ==> false

  myFuncVariable = isNotZero
  fmt.Println(myFuncVariable(9)) // ==> true
}
```

### IIFE

You can immediately call an unnamed function, similar to working with an IIFE in
JS.

```go
var age = 9

var isOlderThanTen bool = func() bool {
  return  age > 10
}()

fmt.Println(isOlderThanTen)
```