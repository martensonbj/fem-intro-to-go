# Types

Go is a _statically typed_ language. Unlike JavaScript, variables must have a
specified type and cannot change once they have been declared. If there is a
discrepancy in types, your program won't compile and you'll get an error.

Let's talk about Go's built in types.

## Integer: `1`, `2`, `44`

- Example: `var age int = 21`
- `int`,`uint8`,`uint16`, `uint32`, `uint64`, `int8`, `int16`, `int32`, `int64`
- How specific you get depends on what type of processor you are working with (won't get into that
  now)
- The `u` indicates an "unsigned" integer, which can only contain positive numbers or zero
- `byte`: alias for `uint8`
- `rune`: alias for `int32`

## Float: `1.5`, `3.14`, `2100`

- Example: `var distanceInMiles float64 = 22.7`
- Floats can contain a decimal point
- `float32`, `float64`
- To see a visual of the difference between how Go stores data in these two types, visit [this playground](https://play.golang.org/p/ZqzdCZLfvC)

## String: `"Marilyn Monroe"`

- Example: `var name string = "Marilyn"`
- Note that Go uses double quotes to indicate a string
- Strings contain an immutable series of bytes that can be accessed by index
  - `fmt.Println(name[0])`
  - `fmt.Println(string(name[0]))`
- Similar to `substr` in JS, you can grab a substring of a string using
  `name[1:3]` ==> `"ar"`
  - This tells Go to start BEFORE index 1, and stop BEFORE index 3.

## Booleans: `true`, `false`

- Example: `var isLessThan bool = 1 < 5`
- Can be generated using the expected logical and comparison operators
- `&&`, `||`, `!`, `<, <=, >, >=, ==, !=`
- Note: Go does not default "falsey" values to a boolean (like `0` or `""`)

## Error: `error`

- Example: `error.Error()` ==> string
- `log.Fatal(err)` will print the error message and stop execution

## TypeOf

- To check the type of a variable, you can use the `TypeOf()` method from the package `reflect`.

```go
  import "fmt"
  import "reflect"

  func main() {
	  var x = "What am I"
	  fmt.Println(reflect.TypeOf(x))
  }
```

- You can also use string formatting within `fmt` to print out a type:
- Note that `%v` is a generic verb for "value", `%T` is a verb for Type.

```go
	var x = "What am I"
	fmt.Printf("The type of `%v` is %T", x, x)
```

## Type Conversion

Sometimes you'll want to convert one type to another. Go lets you do this using the type name as a function.

```go
var age int = 21
var floatAge = float64(age)
fmt.Println(reflect.TypeOf(age))
fmt.Println(reflect.TypeOf(floatAge))
```

Note: Certain types, like strings, need an additional library (`strconv`) to be converted.

