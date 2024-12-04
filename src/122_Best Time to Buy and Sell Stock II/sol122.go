package sol122

//we find the lowest price to buy in and sell it at the highest price
//step1: we select the day when stock has price valley and buy in.
//step2: we select the day when stock has price peak and sell out.
//step3: repeat step1 and step2

func maxProfit(prices []int) int {
	var profit, dayBuyin, daySellout int
	var holdStock bool

	for i := 0; i <= len(prices)-1; i++ {
		if holdStock {
			peak1 := ((i >= 1) && (i <= len(prices)-2) && prices[i] >= prices[i-1] && prices[i] > prices[i+1])
			peak2 := (i == 0 && i+1 <= len(prices)-1 && prices[i] > prices[i+1])
			peak3 := (i == len(prices)-1 && prices[i] > prices[i-1])
			if peak1 || peak2 || peak3 {
				daySellout = i
				profit += (prices[daySellout] - prices[dayBuyin])
				holdStock = false
			}
		} else {
			valley1 := ((i >= 1) && (i <= len(prices)-2) && prices[i] <= prices[i-1] && prices[i] < prices[i+1])
			valley2 := (i == 0 && i+1 <= len(prices)-1 && prices[i] < prices[i+1])
			if valley1 || valley2 {
				dayBuyin = i
				holdStock = true
			}
		}
	}
	return profit
}

func main() {

	test := []int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0}
	res := maxProfit(test)
	println(res)
}
