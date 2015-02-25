## Race Condition Detection in Go

##### Context: My Task Runner

```Go
type TaskData struct {
  In    <-chan interface{}
  Out   chan<- interface{}
  Error chan<- error
}

type TaskRunner interface {
  Run(data *TaskData)
}

func RunTask(data *TaskData, runner TaskRunner) {
  go func() {
    defer close(data.Out)
    runner.Run(data)
  }()
}
```

Note: 
