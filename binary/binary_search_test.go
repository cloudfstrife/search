package binary

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type ints []int

// Compare implement Comparable interface
func (c ints) Compare(v interface{}, i int) int {
	a := v.(int)
	if a > c[i] {
		return 1
	} else if a < c[i] {
		return -1
	} else {
		return 0
	}
}

func TestBinarySearch(t *testing.T) {
	testcases := map[string]struct {
		target []int
		key    int
		asc    bool
		expect int
	}{
		"asc-odd-normal":     {key: 1, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6}, expect: 1},
		"asc-odd-first":      {key: 0, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6}, expect: 0},
		"asc-odd-last":       {key: 6, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6}, expect: 6},
		"asc-odd-mid":        {key: 3, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6}, expect: 3},
		"asc-odd-not-exist":  {key: 9, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6}, expect: -1},
		"asc-even-normal":    {key: 1, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6, 7}, expect: 1},
		"asc-even-first":     {key: 0, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6, 7}, expect: 0},
		"asc-even-last":      {key: 7, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6, 7}, expect: 7},
		"asc-even-mid-1":     {key: 3, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6, 7}, expect: 3},
		"asc-even-mid-2":     {key: 4, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6, 7}, expect: 4},
		"asc-even-not-exist": {key: 9, asc: true, target: []int{0, 1, 2, 3, 4, 5, 6, 7}, expect: -1},

		"desc-odd-normal":            {key: 1, asc: false, target: []int{6, 5, 4, 3, 2, 1, 0}, expect: 5},
		"desc-odd-first":             {key: 6, asc: false, target: []int{6, 5, 4, 3, 2, 1, 0}, expect: 0},
		"desc-odd-last":              {key: 0, asc: false, target: []int{6, 5, 4, 3, 2, 1, 0}, expect: 6},
		"desc-odd-mid":               {key: 3, asc: false, target: []int{6, 5, 4, 3, 2, 1, 0}, expect: 3},
		"desc-odd-not-exist":         {key: 9, asc: false, target: []int{6, 5, 4, 3, 2, 1, 0}, expect: -1},
		"desc-even-normal":           {key: 1, asc: false, target: []int{7, 6, 5, 4, 3, 2, 1, 0}, expect: 6},
		"desc-even-first":            {key: 7, asc: false, target: []int{7, 6, 5, 4, 3, 2, 1, 0}, expect: 0},
		"desc-even-last":             {key: 0, asc: false, target: []int{7, 6, 5, 4, 3, 2, 1, 0}, expect: 7},
		"desc-even-mid-1":            {key: 3, asc: false, target: []int{7, 6, 5, 4, 3, 2, 1, 0}, expect: 4},
		"desc-even-mid-2":            {key: 4, asc: false, target: []int{7, 6, 5, 4, 3, 2, 1, 0}, expect: 3},
		"desc-even-not-exist":        {key: 9, asc: false, target: []int{7, 6, 5, 4, 3, 2, 1, 0}, expect: -1},
		"little-odd":                 {key: 1, asc: true, target: []int{1}, expect: 0},
		"little-odd-not-exist":       {key: 9, asc: true, target: []int{1}, expect: -1},
		"asc-little-even":            {key: 1, asc: true, target: []int{1, 2}, expect: 0},
		"asc-little-even-not-exist":  {key: 9, asc: true, target: []int{1, 2}, expect: -1},
		"desc-little-even":           {key: 1, asc: false, target: []int{2, 1}, expect: 1},
		"desc-little-even-not-exist": {key: 9, asc: false, target: []int{2, 1}, expect: -1},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			got := Search(tc.key, ints(tc.target), 0, len(tc.target)-1, tc.asc)
			if !cmp.Equal(tc.expect, got) {
				t.Errorf("unintended : %s ", cmp.Diff(tc.expect, got))
			}
		})
	}
}
