## Testing in Go

```Go
// lottery.go
package lottery

func ProduceSixLotteryNumbers() []int {
   return []int{4,8,15,16,23,42}
}
```

```Go
// lottery_test.go
package lottery

func TestProduceSixLotteryNumbers(t *testing.T) {
    //...
}
```

Note:
1. 
