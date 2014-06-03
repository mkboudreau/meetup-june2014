
## Test Frameworks & Libraries in Go

##### Ginkgo Test Framework & Gomega Matchers

```Go
func TestExamplePackage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Package Suite")
}
```

```Go
var _ = Describe("some description.... ", func() {
	Context("Some context....", func() {
		It("Should work with table-driven tests", func() {
			for _, testcase := range testcases {
				actual := MyFunction(testcase.input)

				Expect(actual).ToNot(BeNil())
				Expect(actual).To(Equal(testcase.expected))
			}
		})
		It("Should be nil with the text inconceivable", func() {
			badTestCase := "inconceivable"
			actual := MyFunction(badTestCase)
			Expect(actual).To(BeNil())
		})
	})
})
```

Note:
1. 
