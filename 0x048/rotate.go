package main

func rotate1(matrix [][]int) {
	//水平反转
	n := len(matrix)
	for i:= 0;i<n/2;i++{
		for j :=0 ; j<n;j++{
			matrix[i][j],matrix[n-i-1][j] = matrix[n-i-1][j] ,matrix[i][j]
		}
	}

	//主对角线反转
	for i:= 0;i<n;i++{
		for j :=0 ; j<i;j++{
			matrix[i][j],matrix[j][i] = matrix[j][i] ,matrix[i][j]
		}
	}

}
