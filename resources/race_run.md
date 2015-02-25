
## Race Condition Detection in Go

##### Run Test

```bash
$ go test -race
```

Note:
1. github.com/mkboudreau/meetup-june2014/code/taskrunner
2. go test -race
3. ----------------
4. edit taskrunner_string.go (comment defer close)<
5. go test -race