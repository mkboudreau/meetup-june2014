## Race Output

##### WRITE

```
Write by goroutine 5:
  runtime.closechan()
  .../code/taskrunner.func 001()
      .../code/taskrunner/taskrunner.go:22 +0xba
```

##### READ

```
Previous read by goroutine 7:
  runtime.chansend()
  .../code/taskrunner.adaptStringChannelToInterfaceChannel()
      .../code/taskrunner/taskrunner_string.go:47 +0xbd
```
