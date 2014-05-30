## Benchmarking in Go

##### Test File

```Go
// lottery_test.go
package lottery

import (
  "testing"
)

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//
	}
}
```

Note:
1. 
