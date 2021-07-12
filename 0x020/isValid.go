package _x020

func IsValid(s string) bool {
	n := len(s)
	if n ==0 || n == 1{
		return false
	}

	contrast := make(map[uint8]uint8)
	contrast[123]=125//{}
	contrast[91]=93//[]
	contrast[40]=41//()

	var Stack = make([]uint8,n)
	top := 0
	for i,_	 := range s{
		if top == 0 {
			Stack[top] = s[i]
			top++
		}else if contrast[Stack[top-1]] == s[i] {
			top--
		}else {
			Stack[top] = s[i]
			top++
		}
	}

	if top != 0{
		return false
	}
	return true
}