package format

import (
	"strings"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	var tests = []struct {
		input interface{}
		want  string
	}{
		{1, "1"},
		{1 * time.Nanosecond, "1"},
		{[]int64{1}, "[]int64 0x"},
		{[]time.Duration{1 * time.Nanosecond}, "[]time.Duration 0x"},
	}

	for _, test := range tests {
		if res := Any(test.input); !strings.Contains(res, test.want) {
			t.Errorf("Any(%d) = %s, want %s", test.input, res, test.want)
		}
	}
}
