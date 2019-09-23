# Packages

Packages are directories with one or more Go source files that allow Go to reuse code across a program and across files.

Every Go file must belong to a package.

So far, the packages we've seen are `main`, the package we wrote, and `fmt` and `reflect`, which are built into the Go source code. There are [tons of packages](https://golang.org/pkg/) that come out of the box with Go.

To import packages, you list them at the top of your file either like this:

```go
import "fmt"
import "math"
import "reflect"
```

Or more commonly, like this:

```go
import (
  "fmt"
  "math"
  "reflect"
)
```

Some of the packages that ship with Go are:

- `strings` - simple functions to manipulate strings

  - Example Methods: `Contains`, `Count`, `Index`, `HasPrefix`

- `math` - mathematical operations

  - `Pow`, `Abs`, `Ceil`, `Floor` etc
  - `isNaN`, `NaN`

- `io` - handles input/output methods related to the `os` package (like reading
  files)

  - Example Methods: `Copy`, `Reader`, `Writer`

- `os` - methods around operating system functionality

  - Example Methods: `Open`, `Rename`, `CreateFile`

- `testing` - Go's build in test suite

  - Example Methods: `Skip`, `Run`, `Error`

- `net/http` - provides http client and server implementations
  - Example Methods: `Get`, `Post`, `Handle`

## Exported vs Unexported Names

When you import a package, you can only access that package's public, exported names. In Go, these must start with a capital letter.

Anything that starts with a lowercase letter is a private method and will NOT be exported. These are only
visible within the same package.

```go
fmt.Println()
```

Within the `fmt` package, the function `Println` is exported (starts with a
capital letter) and therefore can be accessed within our `main` package.

### Custom Packages

Until now we've been writing all of our go code in the `main` package and only
using functionality from Go's imported libraries.

To demonstrate how packages work, let's create a folder called `utils`, and within that create a file called
`math.go` where we will add a few small functions.

In `math.go` add the following:

```go
package utils

import "fmt"

func printNum(num int) {
	fmt.Println("Current Number: ", num)
}

func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		printNum(v)
		total += v
	}
	return total
}
```

And then in our `packages.go` file, lets import and use this package.

Note that the import name is a path that is relative to the `src` directory.
Often, this will be a github location as your files will live in a github
repository.

Also note that the file name itself isn't necessary as long as that file lives
within the package directory and has its package declaration at the top of the
file.

```go
import {
  "fmt"
  "path/from/src/to/utils" // => ie: fem-intro-to-go/utils
}

func calculateImportantData() int {
	totalValue := utils.Add(1, 2, 3, 4, 5)
	return totalValue
}

func main() {
	total := calculateImportantData()
	fmt.Println(total)
}
```
