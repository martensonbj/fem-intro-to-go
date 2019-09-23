# Arrays

Arrays in Go have some significant differences compared to those in JavaScript.

Let's compare notes.

**JavaScript:**

```javascript
// Initialize an empty array
const grabBag = []
// Eventually this array could have values that represent these types:
const grabBag = [string, int, boolean, float64]
// or
const grabBag = [string, string, string, string, integer, boolean, float64]
```

**Go:**

```go
// Initialize an empty array
var scores [5]float64
// Eventually this array can ONLY contain floats and a length of 5:
[float64, float64, float64, float64, float64]
```

## Defining An Array

In JS, defining an array is pretty chill.

You can make it whatever length you want, you can modify that length, and you can throw any combination of data
types into it.

Go, however, has _opinions_. An array type definition must include a specified length of elements, all of which must be of the same type. In the above example, our variable `scores` is an array of 5 elements, each of which is a float.

Note: An array's length is fixed with the exact length included in its type
definition. This means that `[5]float64` and `[6]float64` are different,
distinct, unequal types.

> _TRY IT_
> In the Go playground, print out the variable `scores` as is. What do you see?

To add elements into this array, both JavaScript and Go allow you to insert elements into a specific index location.

```javascript
grabBag[0] = 'hello' // => ["hello"]
```

```go
ages[0] = 5 // => [5, 0, 0, 0, 0]
```

To add multiple elements, you could start by doing something like this:

```go
  var scores [5]float64
  scores[0] = 9
  scores[1] = 1.5
  scores[2] = 4.5
  scores[3] = 7
  scores[4] = 8
```

Or, to avoid setting each index individually, we can use a _composite literal)_
which containes the type declaration, and a set of initial values:

```go
scores := [5]float64{9, 1.5, 4.5, 7, 8}
```

Going one step further, Go will infer the length of an array when you provide an
initial set of values, and the length can be replaced with an ellipsis.

```go
scores := [...]float64{9, 1.5, 4.5, 7, 8}
```

> _TRY IT_
> Using the array above and a for loop, print out the average score within the array.
> Hint: To find the length of an array if you don't know it, use `len(array)`.

What unexpected errors did you run into? How did you fix them?

## Revisiting Range

When looping over arrays, we can implement a different version of the basic for loop using the `range` keyword:

```go
    var total float64
    for _, value := range scores {
        total += value
    }
```

Let's break it down:

- `i` Still represents the index position we are pointing to in our array
- `value` represents the value of the element at that index (aka `scores[i]`)
- `range` is a keyword that is followed by the variable representing the array we are looping over

> _TRY IT_
> Update the previous for loop from finding the average scores to use `range`.

As we've seen, Go is very particular about types and making sure any declared variable is used within a program. In this case, we aren't using the `i` variable representing the index of the array and we get this error:

```bash
i declared and not used
```

Go uses a single underscore, `_`, to replace an unused variable within a program. Make this modification and watch the program pass.

Defining an array with a specific length can cause some obvious problems.

Adding or removing an element would require redefining the entire variable signature to specify a different length, which often is unknown.

This brings us to Slices.
