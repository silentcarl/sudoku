package main

import (
	"fmt"
	"sort"
	"strconv"
)

type sudoku struct {
	cells [9][9]int
	maybe [9][9][9]bool
}

func (s *sudoku) valid() bool {
	list := [9]int{}
	// 按行
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			list[j] = s.cells[i][j]
		}
		if !isFromOneToNine(list) {
			return false
		}
	}

	// 按列
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			list[j] = s.cells[j][i]
		}
		if !isFromOneToNine(list) {
			return false
		}
	}

	// 按方格
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			list[0] = s.cells[i*3][j*3]
			list[1] = s.cells[i*3][j*3+1]
			list[2] = s.cells[i*3][j*3+2]
			list[3] = s.cells[i*3+1][j*3]
			list[4] = s.cells[i*3+1][j*3+1]
			list[5] = s.cells[i*3+1][j*3+2]
			list[6] = s.cells[i*3+2][j*3]
			list[7] = s.cells[i*3+2][j*3+1]
			list[8] = s.cells[i*3+2][j*3+2]
		}
		if !isFromOneToNine(list) {
			return false
		}
	}

	return true
}

func (s *sudoku) clone() *sudoku {
	c := &sudoku{}
	c.cells = s.cells
	c.maybe = s.maybe
	return c
}

func (s *sudoku) equal(another *sudoku) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.cells[i][j] != another.cells[i][j] {
				return false
			}
		}
	}
	return true
}

func (s *sudoku) init() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				s.maybe[i][j][k] = true
			}
		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.cells[i][j] != 0 {
				s.clearMaybe(i, j, s.cells[i][j])
			}
		}
	}
}

func (s *sudoku) setCell(i, j, val int) {
	s.cells[i][j] = val
	s.clearMaybe(i, j, val)
}

func (s *sudoku) clearMaybe(i, j, val int) {
	// 清理同行、同列的maybe里val
	for k := 0; k < 9; k++ {
		s.maybe[i][k][val-1] = false
		s.maybe[k][j][val-1] = false
	}

	// 清理九格里maybe的val
	col := i - i%3
	row := j - j%3
	for i1 := 0; i1 < 3; i1++ {
		for j1 := 0; j1 < 3; j1++ {
			s.maybe[col+i1][row+j1][val-1] = false
		}
	}

	// 清理本身的maybe
	for k := 0; k < 9; k++ {
		if val-1 == k {
			s.maybe[i][j][k] = true
		} else {
			s.maybe[i][j][k] = false
		}
	}
}

func (s *sudoku) findMaybeOnlyOne() (bool, int, int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.cells[i][j] != 0 {
				continue
			}
			numOfMaybe := 0
			maybe := 0
			for k := 0; k < 9; k++ {
				if s.maybe[i][j][k] {
					numOfMaybe++
					maybe = k + 1
				}
				if numOfMaybe > 1 {
					break
				}
			}
			if numOfMaybe == 1 {
				return true, i, j, maybe
			}
		}
	}
	return false, 0, 0, 0
}

func (s *sudoku) valNotInOtherMaybe(i, j, val int) bool {
	// 同一行其他格子里的maybe有val
	match := true
	for x := 0; x < 9; x++ {
		if x != i && s.maybe[x][j][val-1] {
			match = false
			break
		}
	}
	if match {
		return true
	}

	// 同一列其他格子里的maybe有val
	match = true
	for x := 0; x < 9; x++ {
		if x != j && s.maybe[i][x][val-1] {
			match = false
			break
		}
	}
	if match {
		return true
	}

	// 同一九格其他格子里的maybe有val
	match = true
	col := i - i%3
	row := j - j%3
	for i1 := 0; i1 < 3; i1++ {
		for j1 := 0; j1 < 3; j1++ {
			if i1 == i && j1 == j {
				continue
			}
			if s.maybe[col+i1][row+j1][val-1] {
				match = false
				break
			}
		}
	}
	return match
}

func (s *sudoku) findMaybeMustIt() (bool, int, int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.cells[i][j] != 0 {
				continue
			}
			for k := 0; k < 9; k++ {
				if s.maybe[i][j][k] {
					if s.valNotInOtherMaybe(i, j, k+1) {
						return true, i, j, k + 1
					}
				}
			}
		}
	}
	return false, 0, 0, 0
}

func (s *sudoku) display() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(s.cells[i][j])
			fmt.Print("\t")
		}
		fmt.Println("")
	}
}

func (s *sudoku) displayAll() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			line := []byte{}
			line = strconv.AppendInt(line, int64(s.cells[i][j]), 10)

			if s.cells[i][j] == 0 {
				line = append(line, '(')
				for k := 0; k < 9; k++ {
					if s.maybe[i][j][k] {
						line = strconv.AppendInt(line, int64(k+1), 10)
					} else {
						line = append(line, ' ')
					}
				}
				line = append(line, ')')
			} else {
				for k := 0; k < 9+2; k++ {
					line = append(line, ' ')
				}
			}
			fmt.Print(string(line))
			fmt.Print("\t")
		}
		fmt.Println("")
	}
}

func (s *sudoku) selecteOneMaybe() (bool, int, int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.cells[i][j] != 0 {
				continue
			}
			for k := 0; k < 9; k++ {
				if s.maybe[i][j][k] {
					return true, i, j, k + 1
				}
			}
		}
	}
	return false, 0, 0, 0
}

func (s *sudoku) guess() bool {
	s.init()
	for {
		// 找到只有一个可能的，就是它了
		find, i, j, val := s.findMaybeOnlyOne()
		if find {
			s.setCell(i, j, val)
			continue
		}
		// 找到maybe列表中
		// 例如：i,j格子的maybe列表有 1，4，5，6，但是该格子行或者列或者9格里其他的格子没有maybe是1的，那么这个格子就是1
		find, i, j, val = s.findMaybeMustIt()
		if find {
			s.setCell(i, j, val)
			continue
		}

		copyS := s.clone()
		find, i, j, val = copyS.selecteOneMaybe()
		if find {
			copyS.setCell(i, j, val)
			if copyS.guess() {
				s.cells = copyS.cells
				s.maybe = copyS.maybe
				break
			} else {
				s.maybe[i][j][val-1] = false
			}
			continue
		}
		break
	}
	return s.valid()
}

var validList = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func isFromOneToNine(list [9]int) bool {
	copyList := list[0:9]
	sort.Ints(copyList)
	for i := range validList {
		if copyList[i] != validList[i] {
			return false
		}
	}
	return true
}

func main() {
	// expect := &sudoku{
	// 	cells: [9][9]int{
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 		[9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	// 	},
	// }

	a := &sudoku{
		cells: [9][9]int{
			[9]int{0, 0, 0, 0, 0, 3, 0, 1, 0},
			[9]int{4, 0, 6, 0, 0, 0, 0, 9, 0},
			[9]int{8, 0, 3, 9, 0, 0, 0, 0, 5},
			[9]int{5, 0, 8, 0, 0, 0, 0, 3, 0},
			[9]int{9, 0, 0, 0, 8, 6, 0, 2, 0},
			[9]int{0, 0, 4, 0, 0, 0, 8, 0, 0},
			[9]int{0, 0, 5, 8, 0, 2, 0, 0, 4},
			[9]int{2, 0, 0, 1, 0, 0, 0, 0, 0},
			[9]int{3, 0, 0, 0, 0, 0, 0, 0, 2},
		},
	}
	a.guess()

	if a.valid() {
		a.display()
		fmt.Println("ok")
	} else {
		a.displayAll()
		fmt.Println("error")
	}
}
