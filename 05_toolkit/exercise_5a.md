# Exercise 5a: Testing the Add Method

## Goals
 
- Write a unit test
- Use the TDD method to implement code to make a test pass

## Setup 

- Reference the `utils/math.go` file that contains our previously written `Add` function.
- Create a `strings_test.go` file in the utils directory, and copy the test provided at the bottom of this file (this will be used for Part 2)

## Directions

### Part 1: Write A Test

1. In the previous section we build a custom `utils` package, that had an `Add` method.

2. Add an appropriately named test file, and write a test for this function.

3. Run the test with `go test` and make sure it's passing.

### Part 2: Use TDD to Write A Function

1. In the utils directory navigate to your newly created `strings_test.go` file.

2. Copy the following code into that test file, and watch the test fail.

3. Implement the necessary code to make this test pass.

```go
// strings_test.go
package utils

import "testing"

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := MakeExcited("omg so exciting")
	if actual != expected {
		t.Errorf("Average was incorrect! Expected: %s, Actual: %s", expected, actual)
	}
}
```