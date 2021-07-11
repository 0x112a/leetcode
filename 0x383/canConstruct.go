package _x383

func canConstruct(ransomNote string, magazine string) bool {
	rHash := make(map[string]int)
	mHash := make(map[string]int)
	for _,v :=  range ransomNote{
		rHash[string(v)]++
	}
	for _,v :=  range magazine{
		mHash[string(v)]++
	}

	for k,_ := range rHash{
		if _,ok := mHash[k]; !ok && rHash[k] > mHash[k] {
			return false
		}
	}
	return true
}