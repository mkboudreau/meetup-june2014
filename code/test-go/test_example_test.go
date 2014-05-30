package example

import (
	"fmt"
)

// High-level overview for the entire example package
func Example() {
	result := NewWordEveryNthChar("Hello World", 5)
	fmt.Println(result)
	// Output: H d
}

func ExampleNewWordEveryNthChar_lengthTwo() {
	fmt.Println(NewWordEveryNthChar("Hello", 2))
	// Output: Hlo
}

func ExampleNewWordEveryNthChar_lengthZero() {
	fmt.Println(NewWordEveryNthChar("Hello", 0))
	// Output:
}
func ExampleNewWordEveryNthChar_lengthOne() {
	fmt.Println(NewWordEveryNthChar("Hello", 1))
	// Output: Hello
}

// Negative One is basically a reverse function
func ExampleNewWordEveryNthChar_lengthNegativeOne() {
	fmt.Println(NewWordEveryNthChar("Hello", -1))
	// Output: olleH
}

// This example demonstrates a functional working unit test that is fairly useless as an example since it refers to test data variables that do not get printed
func ExampleNewWordEveryNthChar_wrongWayToDoExamples() {
	testcase := testdata[0]
	fmt.Println(NewWordEveryNthChar(testcase.word, testcase.count))
	// Output: hello world
}
