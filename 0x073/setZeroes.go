package _x073

func setZeroes(matrix [][]int)  {
	row := make(map[int]int)
	column := make(map[int]int)
	r := len(matrix)
	//c := len(matrix[0])

	for i,_ := range matrix{
		for j,_ := range matrix[i]{
			if matrix[i][j] == 0 {
				row[i]++
				column[j]++
			}
		}
	}
	for k,_ := range row{
		for i,_ := range matrix[k] {
			matrix[k][i]=0
		}
	}
	for k,_ := range column{
		for j := 0 ; j < r ; j++{
			matrix[j][k]=0
		}
	}
}