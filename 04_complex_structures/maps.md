# Maps

Maps look and behave similarly to Objects in JavaScript, using a set of key
value pairs.

A map starts with the keyword `map`, followed by the key type, and the
value type.

```go
 var userEmails map[int]string
```

Here we are telling Go to create a map called `userEmails` which will have keys as
integers (think IDs), and values as strings that will represent each email
address.

Let's build one with data in it:

```go
  var userEmails map[int]string
  userEmails[1] = "user1@gmail.com"
  userEmails[2] = "user2@gmail.com"
  fmt.Println(userEmails)
```

> _TRY IT_
> Throw that code into the Go playground. What happens? What are we missing?

If we run the above code as-is, we'll see the following error:

```go
panic: assignment to entry in nil map
```

The (run-time!) error here, indicated by the keyword `panic`, is telling us that
we're trying to assign entries to a map that doesn't exist yet.

Look at where we are declaring our `userEmails` variable. Here, we are telling
Go to create a variable called `userEmails` that will have a type of map with
particular key value pair types. We haven't actually set that variable to a
value of any kind, nor have we allocated any memory where it can be stored.

Remember the `make` function from our slice examples? Let's do that now.

```go
  var userEmails map[int]string
  userEmails = make(map[int]string)

  userEmails[1] = "user1@gmail.com"
  userEmails[2] = "user2@gmail.com"
  fmt.Println(userEmails)
```

Or, with the shorthand syntax:

```go
  userEmails := make(map[int]string)
  userEmails[1] = "user1@gmail.com"
  userEmails[2] = "user2@gmail.com"
  fmt.Println(userEmails)
```

Now we should see:
`map[1:user1@gmail.com 2:user2@gmail.com]`

As with arrays, we can simplify creating our map using abbreviated syntax. Note
that here we don't need the `make` keyword since we are providing Go with an
initial amount of data to store in memory:

```go
userEmails := map[int]string{
  1: "user1@gmail.com",
  2: "user2@gmail.com",
}
fmt.Println(userEmails)
```

## Mutating and Reading Maps

To ask for a specific value from a map, you'd use syntax similar to interacting
with objects in JS:

```go
userEmails[1] // ==> user1@gmail.com
```

To modify an element in an array, simply reassign its value:

```go
userEmails[1] = "newUser1@gmail.com"
fmt.Println(userEmails)
```

> _TRY IT_
> What happens if you ask for a key that doesn't exist? Try asking for
> `userEmails[3]`.

Go has a clever way to handle checking for null values. Looking up a value
within a map actually returns two elements: the value (if it exists), and a
boolean for whether or not that value existed.

```go
  firstEmail, ok := userEmails[1]
  missingEmail, ok := userEmails[3]
  fmt.Println("Email:", firstEmail, "Present?", ok)
  fmt.Println("Email", missingEmail, "Present?", ok)
```

This makes it easy to add logic around checking for null values within our
program using an if statement:

```go
  if email, ok :=  userEmails[1]; ok {
    fmt.Println(email)
  } else {
    fmt.Println("I don't know what you want from me")
  }
```

First, we create the two variables `email` and `ok` inside our if
statement and set them to the two elements we get back from a lookup in our map.

Then, we check what boolean we got back in our `ok` variable. The code within
the first block will fire only if `ok` returns true, otherwise our if block will
continue to the else block.

### Deleting

```go
delete(userEmails, 2)
```

### Iterating

We can iterate over a map by using the familiar `range` keyword, however now our
for loop will be iterating over the `key` and `value` variables.

```go
for k, v := range userEmails {
    fmt.Printf("%s has an ID of %d.\n", v, k)
}
```

