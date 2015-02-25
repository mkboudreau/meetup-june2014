## Race Output

##### Write by goroutine 5: runtime.closechan()

```
  .../code/taskrunner.func 001()
      /Users/.../code/taskrunner/taskrunner.go:22 +0xba
```
```Go
func RunTask(data *TaskData, runner TaskRunner) {
	go func() {
		defer close(data.Out)
		runner.Run(data)
	}() // line 22
}
```

##### Previous read by goroutine 7: runtime.chansend()

```
  .../code/taskrunner.adaptStringChannelToInterfaceChannel()
      /Users/.../code/taskrunner/taskrunner_string.go:47 +0xbd
```
```Go
func adaptStringChannelToInterfaceChannel(in <-chan string, out chan<- interface{}, err chan<- error) {
	defer close(out)

	for someString := range in {
		out <- someString // line 47
	}
}
```
