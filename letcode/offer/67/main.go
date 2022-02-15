// @Description _7
// @Author caopengfei 2021/5/7 18:48
package main

import (
	"github.com/patrickmn/go-cache"
	"math"
	"time"
)

func main() {
	strToInt("2147483648")
	c := cache.New(30*time.Second, 10*time.Second)


	c.Set("Title", "Spring Festival", cache.DefaultExpiration)
}

func strToInt(str string) int {
	res := 0
	flag := 1
	sign := false
	for i := 0; i < len(str); i++ {
		if str[i] == ' '&&!sign {
			continue
		}
		if str[i] == '+' {
			if(sign) {
				return 0
			}
			sign = true
			continue
		}
		if str[i] == '-' {
			if(sign) {
				return 0
			}
			flag *= -1
			sign = true
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			break
		}
		if res > math.MaxInt32 || res == math.MaxInt32 && str[i] >= math.MaxInt32%10 {
			if flag > 0 {
				return math.MaxInt32 -1
			}
			return (math.MaxInt32+1) * flag
		}else {
			res = res * 10 + int(str[i]) - int('0')
		}
	}
	return res * flag
}