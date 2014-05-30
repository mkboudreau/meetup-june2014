package lottery

import (
	"testing"
)

func TestProduceSixLotteryNumbers(t *testing.T) {
	lottery := ProduceSixLotteryNumbers()

	if lottery == nil {
		t.Errorf("lottery should not be nil")
	} else if len(lottery) != 6 {
		t.Errorf("expected lottery to be size 6, but was %v", len(lottery))
	}
}

func BenchmarkProduceSixLotteryNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProduceSixLotteryNumbers()
	}
}
