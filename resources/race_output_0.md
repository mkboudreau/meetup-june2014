## Race Output

```
==================
WARNING: DATA RACE
Write by goroutine 5:
  runtime.closechan()
      /usr/local/go/src/pkg/runtime/chan.c:1232 +0x0
  github.com/mkboudreau/meetup-june2014/code/taskrunner.func·001()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner.go:22 +0xba

Previous read by goroutine 7:
  runtime.chansend()
      /usr/local/go/src/pkg/runtime/chan.c:155 +0x0
  github.com/mkboudreau/meetup-june2014/code/taskrunner.adaptStringChannelToInterfaceChannel()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner_string.go:47 +0xbd
  github.com/mkboudreau/meetup-june2014/code/taskrunner.func·004()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner_string.go:25 +0xb7

Goroutine 5 (running) created at:
  github.com/mkboudreau/meetup-june2014/code/taskrunner.RunTask()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner.go:22 +0xd6
  github.com/mkboudreau/meetup-june2014/code/taskrunner.ChainTasks()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner.go:32 +0x297
  github.com/mkboudreau/meetup-june2014/code/taskrunner.func·005()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner_combo_test.go:13 +0x241
  github.com/smartystreets/goconvey/convey.(*action).Invoke()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/registration.go:42 +0x42
  github.com/smartystreets/goconvey/convey.(*scope).visit()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:50 +0x75
  github.com/smartystreets/goconvey/convey.(*scope).visitChild()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:65 +0xb6
  github.com/smartystreets/goconvey/convey.(*scope).visitChildren()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:60 +0x64
  github.com/smartystreets/goconvey/convey.(*scope).visit()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:51 +0x83
  github.com/smartystreets/goconvey/convey.(*scope).visitChild()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:65 +0xb6
  github.com/smartystreets/goconvey/convey.(*scope).visitChildren()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:60 +0x64
  github.com/smartystreets/goconvey/convey.(*scope).visit()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/scope.go:51 +0x83
  github.com/smartystreets/goconvey/convey.(*runner).Run()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/runner.go:96 +0xbf
  github.com/smartystreets/goconvey/convey.(*suiteContext).Run()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/context.go:44 +0x227
  github.com/smartystreets/goconvey/convey.register()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/doc.go:54 +0x7e
  github.com/smartystreets/goconvey/convey.Convey()
      /Users/michaelboudreau/Development/go/src/github.com/smartystreets/goconvey/convey/doc.go:27 +0x56
  github.com/mkboudreau/meetup-june2014/code/taskrunner.TestChainingMultipleTypesOfTaskRunner()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner_combo_test.go:59 +0x147
  testing.tRunner()
      /usr/local/go/src/pkg/testing/testing.go:391 +0x10f

Goroutine 7 (running) created at:
  github.com/mkboudreau/meetup-june2014/code/taskrunner.StringTaskRunner.Run()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner_string.go:26 +0x20b
  github.com/mkboudreau/meetup-june2014/code/taskrunner.func·001()
      /Users/michaelboudreau/Development/go/src/github.com/mkboudreau/meetup-june2014/code/taskrunner/taskrunner.go:21 +0xb9
==================
.........................................PASS
Found 1 data race(s)
```