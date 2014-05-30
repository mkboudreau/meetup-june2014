## Race Output

##### Write by goroutine 5:

```
Write by goroutine 5:
  runtime.closechan()
      /usr/local/go/src/pkg/runtime/chan.c:1232 +0x0
  .../code/taskrunner.func 001()
      /Users/.../code/taskrunner/taskrunner.go:22 +0xba
```

##### Previous read by goroutine 7:

```
Previous read by goroutine 7:
  runtime.chansend()
      /usr/local/go/src/pkg/runtime/chan.c:155 +0x0
  .../code/taskrunner.adaptStringChannelToInterfaceChannel()
      /Users/.../code/taskrunner/taskrunner_string.go:47 +0xbd
  .../code/taskrunner.func 004()
      /Users/.../code/taskrunner/taskrunner_string.go:25 +0xb7
```
