package _x036

func isValidSudoku(board [][]byte) bool {
	var row  [9]map[int]int
	var columns  [9]map[int]int
	var boxes  [9]map[int]int
	for i :=0; i<9 ; i++{
		row[i] = make(map[int]int)
		columns[i] = make(map[int]int)
		boxes[i] = make(map[int]int)
	}

	for i := 0; i<9 ;i++{
		for j := 0; j < 9; j++ {
			n := int(board[i][j])
			boxIndex := i/3*3 + j/3
			if n == 46{
				continue
			}
			row[i][n]++
			columns[j][n]++
			boxes[boxIndex][n]++
			if row[i][n] > 1 || columns[j][n] >1 || boxes[boxIndex][n] >1{
				return false
			}
		}
	}
	return true
}