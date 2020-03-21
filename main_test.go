package main

import (
	"testing"
)

type testcase struct {
	s      *sudoku
	expect *sudoku
}

func TestSudoku(t *testing.T) {
	testcases := []testcase{
		{
			s: &sudoku{
				cells: [9][9]int{
					[9]int{5, 3, 0, 0, 7, 0, 0, 0, 0},
					[9]int{6, 0, 0, 1, 9, 5, 0, 0, 0},
					[9]int{0, 9, 8, 0, 0, 0, 0, 6, 0},
					[9]int{8, 0, 0, 0, 6, 0, 0, 0, 3},
					[9]int{4, 0, 0, 8, 0, 3, 0, 0, 1},
					[9]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
					[9]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
					[9]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
					[9]int{0, 0, 0, 0, 8, 0, 0, 7, 9},
				},
			},
			expect: &sudoku{
				cells: [9][9]int{
					[9]int{5, 3, 4, 6, 7, 8, 9, 1, 2},
					[9]int{6, 7, 2, 1, 9, 5, 3, 4, 8},
					[9]int{1, 9, 8, 3, 4, 2, 5, 6, 7},
					[9]int{8, 5, 9, 7, 6, 1, 4, 2, 3},
					[9]int{4, 2, 6, 8, 5, 3, 7, 9, 1},
					[9]int{7, 1, 3, 9, 2, 4, 8, 5, 6},
					[9]int{9, 6, 1, 5, 3, 7, 2, 8, 4},
					[9]int{2, 8, 7, 4, 1, 9, 6, 3, 5},
					[9]int{3, 4, 5, 2, 8, 6, 1, 7, 9},
				},
			},
		},
	}
	for i, tc := range testcases {
		tc.s.guess()
		if !tc.s.equal(tc.expect) {
			t.Error(i)
		}
	}

}
