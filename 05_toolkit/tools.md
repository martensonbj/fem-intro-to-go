# Go Tools

## Compiling Go Packages

In the previous Hello World example we used the command `go run ...` to compile and run our
`main.go` program. Let's discuss a few others:

## Build & Install

The two commands `go build` and `go install` will compile all of your go
packages and dependencies and put the executable in the current directory (when run without any flags).

`go install` will also put the compiled dependencies into the `pkg` directory
within your go workspace.

`go build` lets you build an executable file locally, which lets you test your
application without having to deploy it immediately.

## go fmt

`go fmt` will reformat your code based on Go standards. This is often run as a
linter, (like on save of a file). 

## go list 

`go list` will list all of the packages in that current directory.

## go vet

`go vet` runs through your source code and yells about any weird constructs

It catches small syntax errors that often slip through code review or unit tests - like `fmt.Printf` calls that have misaligned arguments. 

## go doc 

`go doc` we mentioned earlier, will give you access to local help docs

## go get

In order to install an external package, Go provides the `go get` command.
Unlike installing an `npm` module, in Go you need to specify the location of
where that executable lives. For example, if we want to use the `golint`
package, we would install it like so:

`go get -u golang.org/x/lint/golint`

## go lint

Go lint is concerned with style mistakes and maintaining a consistent coding style that is miantained at Google.

Install `golint` using the command above.

To find out where the library was put, run `go list -f {{.Target}} golang.org/x/lint/golint`

For `golint` to be used globally, make sure the path you see from the above command is included in your `$PATH` variable.