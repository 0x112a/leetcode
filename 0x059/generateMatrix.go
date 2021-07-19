package _x059

type direct struct {
	x,y int
}

var directions = []direct{{0,1},{1,0},{0,-1},{-1,0}}

func generateMatrix(n int) [][]int {
	ans := make([][]int,n)
	for i,_ := range ans{
		ans[i]= make([]int,n)
	}
	nn := n*n
	var c,r int
	var directionIndex int

	for i:=1;i<=nn;i++{
		ans[r][c] = i

		nextRow ,nextColunm := r + directions[directionIndex].x,c + directions[directionIndex].y

		if nextRow < 0 || nextRow>=n || nextColunm<0 || nextColunm>=n || ans[nextRow][nextColunm] != 0 {
			directionIndex = (directionIndex+1)%4
		}

		r +=directions[directionIndex].x
		c += directions[directionIndex].y

	}
	return ans
}
