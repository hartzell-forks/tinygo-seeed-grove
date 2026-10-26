package binify

import (
	"errors"
)

// Errors
var (
	ErrBadPercent = errors.New("Bad percentage, must be between 0 and 100, inclusive")
)

func Binify(percent float32) (int, error) {
	bin_size := 12.5

	if percent < 0 || percent > 100 {
		return 0, ErrBadPercent
	}

	// Special case, 8 is actually [87.5, 100)
	// But we need to do *something* with 100, so make that interval a
	// bit bigger, [87.5, 100].
	if percent == 100 {
		return 8, nil
	}

	bin := int((percent / float32(bin_size)) + 1)

	return bin, nil
}
