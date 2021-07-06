package _x121

//func maxProfit(prices []int) int {
//	var pMin int
//	for i,v := range prices{
//		if prices[pMin] > v{
//			pMin = i
//		}
//	}
//	if pMin == len(prices)-1{
//		return 0
//	}
//	var Max int
//	for i,n:= pMin,len(prices); i< n ; i++{
//		if prices[i] > Max {
//			Max = prices[i]
//		}
//	}
//	return Max-prices[pMin]
//}

func maxProfit(prices []int) int {
	var pMin,Max int
	for i,v := range prices{
		if v < prices[pMin]{
			pMin = i
		}
		if Max < v - prices[pMin] {
			Max = v-prices[pMin]
		}
	}
	return Max
}
