# Anatomy of a Go file

Let's compare the basic file structures between Go and JavaScript.

```javascript
import MyJSClass from '../MyJSClass'

import React from 'react'

// Access to MyJSClass and React methods and variables are now possible in this
// file

const Component = () => {
  // Code to execute
  console.log('Hello Front End Masters!')
}

export default Component
```

```go
package main // 1

import "fmt" // 2

// Access to main and fmt methods/variables is now possible in this file

func main() { //3
  // Code to execute lives here
  fmt.Println("Hello Front End Masters!"
} // 5
```

1. Package Declaration

  - Packages are Go's way of organizing and reusing code within an application
  - These packages, like classes, expose different variables that can be used within a file
  - Every go program must have at least a `main` package
  - When you run `go install` Go creates a binary file which will call the `main()` function of your program
  - To find details on what a package does from your terminal, for instance `fmt`, you can run `go doc fmt`

2. Imported packages and libraries

  - Exposes code from built in go libraries, third party packages, or internally created packages your program needs
  - In this case, the package `fmt` comes with go and provides formatting methods for standard input and output
  - Also, it's pronounced "fumpt". 🙄

3. The `main()` go function signature

  - This function initializes the go program
  - Starts with a keyword `func`
  - The name of the function - `main`
  - A list of optional parameters - not present
  - An optional return statement - not present
  - Opening curly brace

4. Code to be executed

5. Closing curly brace

  - Indicates the end of a function
  - Is often preceded by a `return` statement, similar to JS
