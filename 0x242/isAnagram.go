package _x242

func isAnagram(s string, t string) bool {
	slen,tlen := len(s),len(t)
	if slen != tlen{
		return false
	}

	sHash,tHash := make(map[string]int),make(map[string]int)
	for _,v := range s{
		sHash[string(v)]++
	}
	for _,v := range t{
		tHash[string(v)]++
	}

	for k,v :=	range sHash{
		if _,ok := tHash[k]; !ok || tHash[k] != v{
			return false
		}
	}
	return true
}