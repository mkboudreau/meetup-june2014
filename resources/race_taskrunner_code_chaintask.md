## Race Condition Detection in Go

##### Context: My Task Runner

```Go
func ChainTasks(data *TaskData, runners ...TaskRunner) {
  for i, runner := range runners {
    if i < len(runners)-1 {
      ch := make(chan interface{})
      RunTask(&TaskData{In: data.In, Out: ch, Error: data.Error}, runner)
      data.In = ch
    } else {
      RunTask(&TaskData{In: data.In, Out: data.Out, Error: data.Error}, runner)
    }
  }
}
```

Note: 
