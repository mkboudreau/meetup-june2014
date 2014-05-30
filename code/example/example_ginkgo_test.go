package example_test

import (
	. "../example"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

func TestExamplePackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Package Suite")
}

var _ = Describe("Example Package", func() {

	Context("NewWordEveryNthChar Function", func() {
		It("Test table-driven permutations", func() {
			for _, testcase := range testdata {
				actual := NewWordEveryNthChar(testcase.word, testcase.count)

				Expect(actual).ToNot(BeNil())
				Expect(actual).To(Equal(testcase.expected))
			}
		})
		It("Should return itself when count = 1", func() {
			testcase := "inconceivable"
			count := 1
			expected := testcase

			actual := NewWordEveryNthChar(testcase, count)

			Expect(actual).ToNot(BeNil())
			Expect(actual).To(Equal(expected))
		})

	})

})
