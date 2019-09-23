# Unit Testing

## Testing a Method

Go includes a package, `testing`, that contains all the functions necessary to
run a test suite and command to run those tests, `go test`.

Test file names must end with `_test.go` and the Go compiler will know to ignore
these unless `go test` is run.

The test files need to be in the same package as the method they are testing.

The test name should start with `Test` followed by the name of the
function you are testing.

The only parameter passed into the test should be `t *testing.T`.

Note: We will discuss what that `*` means shortly.

```go
// average_test.go

package utils

import "testing"

// The test should always accept the same argument, t, of type *testing.T)
func TestAverage(t *testing.T) {
  // the value you expect to get from the tested function
  expected := 4
  actual := utils.average(1, 2, 3)
  if actual != expected {
    t.Errorf("Average was incorrect! Expected: %d, Actual: %d", expected, actual)
  }
}
```

To run tests, use the command `go test` from within the directory where the test
file lives.

_TRY IT_
Test The Add Function from our `math.go` file in utils.
