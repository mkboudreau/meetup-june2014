## Test Frameworks & Libraries in Go

##### GoConvey Test Framework & oglematchers Matchers

```Go
func TestSpec(t *testing.T) {
    Convey("Given some integer with a starting value", t, func() {
        x := 1

        Convey("When the integer is incremented", func() {
            x++

            Convey("The value should be greater by one", func() {
                So(x, ShouldEqual, 2)
            })
        })
    })
}
```

Note:
1. 
