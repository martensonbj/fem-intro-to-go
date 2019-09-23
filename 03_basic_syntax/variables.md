# Variables

Variables in Go are written in camel case, similar to JavaScript, and can be defined a few different ways.

**Option 1:**

- Initialize a variable with a type and the keyword `var`
- Looks similar to JavaScript with an additional type declaration

```go
var name string = "Beyonce"
```

**Option 2:**

- Go can infer the type of initialized variables and the type declaration can be
  omitted.

```go
var name = "Beyonce"
```

**Option 3:**

- Feels similar to using `let` in JavaScript to instantiate and empty variable, but instead of storing the value as `undefined`, Go will default to the zero value for that type
- string: `""`
- int: `0`
- float: `0`
- bool: `false`

```go
var name string
// This will have a default value of "" that can be modified later
```

**Option 4:**

- Declare multiple variables at once

```go
var name1, name2 string = "Beyonce", "Lizzo"
```

**Option 5:**

- Can only be used within a function
- Most commonly used pattern
- To declare and assign a default value at the same time
- You can omit the keyword `var`, and go will infer the value

```go
name := "Beyonce"
```

#### Var vs Const

Consts, like in JavaSript, are variables whose values cannot be changed.
Attempting to modify a const will result in a run-time error.

**Sidebar**

- RunTime vs CompileTime
  - Reminder that compile time errors involve syntax or type-checking. They occur when you, as the developer, are compiling your code.
  - Run time errors happen after the program is compiled and a user or browser is trying to execute that code.
  - Examples of a run time error: Trying to open a file or url that doesn't exist, running out of memory, trying to do something that syntactically is legit but isn't valid in real life (like dividing by 0).
