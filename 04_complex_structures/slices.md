## Make

Before we get into slices, let's talk about a specific Go builtin function called
`make()`.

According to [the docs](https://golang.org/pkg/builtin/#make), the make function allocates space in memory for, and initializes a `slice`, `map`, or `channel`.

This is necessary because, as seen with Arrays, Go is anticipating the need to set
aside a specific length of memory for a variable being created. Since `slice`
and `map` allow you to define a variable _without_ a set length, `make` helps Go
establish where this variable will live and how to access it.

In the example below, we are initializing an array, and a slice.

```go
var myArray [5]int
var mySlice []int

myArray[0] = 1
mySlice[0] = 2

fmt.Println(myArray)
fmt.Println(mySlice)
```

> TRY_IT
> What happens when you print the variables above?

If that was all of the code we had, Go would have no way of knowing how much memory to allocate for the slice, since it has no association with an underlying array (yet). This is where `make` comes in. 

Because slices are _segments of an array_, they must always be associated with an underlying array.

```go
	var myArray [5]int
	// var mySlice []int = make([]int, 5)
	var mySlice []int = make([]int, 5, 10)
	// var mySlice = make([]int, 5)

	myArray[0] = 1
	mySlice[0] = 2

	fmt.Println(myArray)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
```

## Slice

A slice is a segment of an array. Like an array, a slice must contain a single type of element. Elements can still be accessed by their index, a slice must start with an initial length, but unlike arrays their lengths can change.

Let's say we have a fixed array that looks like this:

```go
fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}
```

We could access and create a slice of that by specifying two indices from within
that array, separated by a colon.

In the following example, we are asking for elements starting at index 1, up to but
not including index 3.

```go
var splicedFruit []string = fruit[1:3]  // ==> ["pear", "apple",]
```

// SEE SLIDE 

NOTE: When passing a slice around to various functions, you are essentially
passing the _header_ of the slice, the "summary" of what that slice references,
which includes a hidden pointer to the underlying array.


`fruit[:3]` ==> Everything up to (but not including) index 3
`fruit[3:]` ==> Everything after (and including) index 3

Instantiating a slice (without setting aside any memory for it) looks almost identical to instantiating an array, just without a specific length.

```go
  var scores[]float64
```


### Len, Cap

The associated underlying array can be visualized by using built in functions to see the difference.

```
fruitArray := [5]string{"banana", "pear", "apple", "kumquat", "peach"}
var splicedFruit []string = fruitArray[1:3]  // ==> ["pear", "apple",]
```

_Try It_
What happens when you ask for the `len()` of the array? (Length)
What about the `cap()` (Capacity)

This indicates that the longest this slice can be, is the length of its
underlying array.


What if you want to increase the length of a slice?



### Slice Helper Methods

There are built in helper methods associated with the `slice` type, similar to
prototype methods in JavaScript. The one we will use later in this course is
`append`.

#### Append:

- `append(originalSlice, newEl1, newEl2)`

```go
  slice1 := []int{1, 2, 3}
  slice2 := append(slice1, 4, 5)
  fmt.Println(slice1, slice2)
```

Another interesting method that takes a second to think through coming from the
FE is go's `copy`.

#### Copy:

What do you expect the following `mysteryValue` variable to print out?

```go
	originalSlice := []int{1, 2, 3}
	destination := make([]int, len(originalSlice))
	fmt.Println("Before Copy:", originalSlice, destination)

	mysteryValue := copy(destination, originalSlice)

	fmt.Println("After Copy:", originalSlice, destination, mysteryValue)
```

> _TRY IT_
> Run the example code from the Copy block.
> What do you notice about the value of `mysteryValue`? 

Check out [this blog
post](https://blog.golang.org/go-slices-usage-and-internals) for more great info
on slices and arrays.