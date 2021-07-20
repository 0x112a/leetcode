package _x240

func searchMatrix(matrix [][]int, target int) bool {
	r := len(matrix)-1
	n := len(matrix[0])
	c := 0
	for r >= 0 && c < n{
		if matrix[r][c] > target{
			r--
		}else if matrix[r][c] < target {
			c++
		}else {
			return true
		}
	}
	return false
}
