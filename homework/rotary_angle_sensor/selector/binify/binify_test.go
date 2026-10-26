package binify

import (
	"errors"
	"testing"
)

var binTests = []struct {
	input float32
	bin   int
	err   error
}{
	{0, 1, nil},
	{.1, 1, nil},
	{12.4999, 1, nil},
	{12.5, 2, nil},
	{12.51, 2, nil},
	{24.9, 2, nil},
	{25, 3, nil},
	{99, 8, nil},
	{100, 8, nil},
	// errors
	{-0.1, 0, ErrBadPercent},
}

func TestBinify(t *testing.T) {
	for testNumber, testCase := range binTests {
		bin, err := Binify(testCase.input)
		if !errors.Is(err, testCase.err) {
			t.Errorf("Test number %d: err failed: expected=%v, got=%v", testNumber, testCase.err, err)
		}

		if bin != testCase.bin {
			t.Errorf("Test number %d: bin incorrect: expected=%d, got=%d", testNumber, testCase.bin, bin)
		}
	}
}
