## Race Condition Detection in Go

##### Context: My Task Runner

```Go
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
	in <- "TESTING A"
	in <- "123"
	in <- "HELLO"
	in <- "456"
	in <- "TESTING B"
	close(in)
}()
```

Note: 
