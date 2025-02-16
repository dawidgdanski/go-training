package language

import "testing"

func Test_divAndRemainder_should_return_error_when_denominator_is_zero(t *testing.T) {
	_, _, err := divAndRemainder2(1, 0)

	if err == nil {
		t.Error("divAndRemainder2 should return an error when denominator is set to 0.")
	}
}
