package add

import (
	"fmt"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	sliceToString := func(numbers []int64) string {
		strSlice := []string{}

		for _, num := range numbers {
			strSlice = append(strSlice, fmt.Sprintf("%v", num))
		}

		return strings.Join(strSlice, ",")
	}
	tests := []struct {
		expected int64
		numbers  []int64
	}{
		{15, []int64{1, 2, 3, 4, 5}},
		{0, []int64{}},
		{5, []int64{1, 2, 3, 4, -5}},
		{-6, []int64{-1, -2, -3}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%v)", sliceToString(test.numbers)), func(t *testing.T) {
			result := Add(test.numbers)
			if test.expected != result {
				t.Errorf("Expected: %v, Result: %v", test.expected, result)
			}
		})
	}
}
