package _x118

func generate(numRows int) [][]int {
	reslut := make([][]int,numRows)
	for i:=0; i<numRows; i++{
		reslut[i] = make([]int,i+1)
		reslut[i][0] = 1
		reslut[i][i] = 1
		for j := 1; j<i ;j++{
			reslut[i][j] = reslut[i-1][j] + reslut[i-1][j-1]
		}
	}
	return reslut
}