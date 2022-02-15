// @Description offer063
// @Author caopengfei 2021/5/7 18:17
package main


func main() {
	println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	max := 0
	for _,price := range prices {
		if minPrice > price {
			minPrice = price
		}
		if max < price - minPrice {
			max = price - minPrice
		}
	}
	return max
}

