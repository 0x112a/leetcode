package _x566


func matrixReshape(mat [][]int, r int, c int) [][]int {
	//判断是否能转换
	matr,matc := len(mat),len(mat[0])
	if matr * matc != r*c {
		return mat
	}

	//转换矩阵
	result :=make([][]int,c)
	cur := 0
	for ind, _ := range result {
		result[ind] = make([]int,r)
		for index,_ := range result[ind]{
			//类似位示图
			i := cur/matr
			j := cur%matr
			result[ind][index] = mat[i][j]
		}
	}

	return result

}
