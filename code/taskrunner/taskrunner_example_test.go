package taskrunner

import (
	"fmt"
)

func Example() {
	in := make(chan interface{})
	out := make(chan interface{})
	err := make(chan error)
	taskData := &TaskData{
		In:    in,
		Out:   out,
		Error: err,
	}

	RunTask(taskData, &PassThroughTaskRunner{})
	go func() {
		in <- "TESTING"
		in <- "ABC"
		in <- "123"
		close(in)
	}()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	// Output: TESTING
	// ABC
	// 123
}

func ExampleRunTask() {
	in := make(chan interface{})
	out := make(chan interface{})
	err := make(chan error)
	taskData := &TaskData{
		In:    in,
		Out:   out,
		Error: err,
	}

	RunTask(taskData, &PassThroughTaskRunner{})
	go func() {
		in <- "TESTING"
		in <- "ABC"
		in <- "123"
		close(in)
	}()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	// Output: TESTING
	// ABC
	// 123
}

func ExampleRunTask_stringFilter() {
	in := make(chan interface{})
	out := make(chan interface{})
	err := make(chan error)
	taskData := &TaskData{
		In:    in,
		Out:   out,
		Error: err,
	}

	RunTask(taskData, &FilterString{Filter: "HELLO"})

	go func() {
		in <- "TESTING"
		in <- "HELLO"
		in <- "ABC"
		in <- "123"
		close(in)
	}()

	fmt.Println(<-out)
	<-err
	fmt.Println(<-out)
	fmt.Println(<-out)

	// Output: TESTING
	// ABC
	// 123
}

type PassThroughTaskRunner struct{}

func (runner *PassThroughTaskRunner) String() string {
	return "PassThroughTaskRunner"
}
func (runner *PassThroughTaskRunner) Run(data *TaskData) {
	defer func() {
		if r := recover(); r != nil {
			data.Error <- fmt.Errorf(" (test runner impl) Problem with run on PassThroughTaskRunner: %v", r)
		}
	}()
	for obj := range data.In {
		data.Out <- obj
	}
}
