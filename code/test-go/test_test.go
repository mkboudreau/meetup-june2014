package example

import (
	//"fmt"
	"testing"
)

var testdata = []struct {
	word     string
	count    int
	expected string
}{
	{"hello world", 1, "hello world"},
	{"hello world", 2, "hlowrd"},
	{"hello world", 3, "hlwl"},
	{"hello world", 4, "hor"},
	{"hello world", 0, ""},
	{"hello world", -1, "dlrow olleh"},
	{"hello world", -2, "drwolh"},
	{"hello world", -3, "dooe"},
	{"hello world", -4, "dwl"},
}

func TestNewWordEveryNthChar(t *testing.T) {
	for _, testcase := range testdata {
		actual := NewWordEveryNthChar(testcase.word, testcase.count)
		if actual != testcase.expected {
			t.Errorf("Actual [%v] does not equal Expected [%v]", actual, testcase.expected)
		}
	}
}

//Hello
func BenchmarkNewWordEveryNthChar(b *testing.B) {
	testIndex := 0
	for i := 0; i < b.N; i++ {
		if testIndex >= len(testdata) {
			testIndex = 0
		}
		NewWordEveryNthChar(testdata[testIndex].word, testdata[testIndex].count)
		testIndex++
	}
}
