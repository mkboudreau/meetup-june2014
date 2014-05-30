package example

func NewWordEveryNthChar(original string, n int) string {
	if n > 0 {
		return newWordEveryNthCharGoingForward(original, n)
	} else if n < 0 {
		return newWordEveryNthCharGoingReverse(original, n)
	}
	return ""
}

func newWordEveryNthCharGoingForward(original string, n int) string {
	newString := make([]byte, len(original))
	newStringIndex := 0

	for i := 0; i < len(original); i += n {
		newString[newStringIndex] = original[i]
		newStringIndex++
	}
	return string(newString[:newStringIndex])
}

func newWordEveryNthCharGoingReverse(original string, n int) string {
	newString := make([]byte, len(original))
	newStringIndex := 0

	for i := len(original) - 1; i >= 0; i += n {
		newString[newStringIndex] = original[i]
		newStringIndex++
	}
	return string(newString[:newStringIndex])
}
