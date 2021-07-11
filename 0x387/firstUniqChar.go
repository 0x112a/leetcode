package _x387

func firstUniqChar(s string) int {
	sHash := make(map[string][]int)
	for i,v := range s{
		stemp := string(v)
		if sHash[stemp] == nil{
			sHash[stemp] = make([]int,2)
		}
		sHash[stemp][0] = i
		sHash[stemp][1]++
	}
	ansIndex := len(s)
	for _,v := range sHash{
		if v[1] == 1 && ansIndex > v[0]{
			ansIndex = v[0]
		}
	}

	if ansIndex == len(s){
		return -1
	}

	return ansIndex
}

//func firstUniqChar(s string) int {
//	var cnt [26]int
//
//	for i:=0;i<len(s);i++{
//		cnt[s[i]-'a']++
//	}
//	for i:=0;i<len(s);i++{
//		if cnt[s[i]-'a']==1{
//			return i
//		}
//	}
//	return -1
//}