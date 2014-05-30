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
    lottery := ProduceSixLotteryNumbers()
    
    if lottery == nil {
        t.Errorf("lottery should not be nil")
    } else if len(lottery) != 6 {
        t.Errorf("expected lottery to be size 6, but was %v", len(lottery))
    }
}
```

Note:
1. 
