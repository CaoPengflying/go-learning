// @Description changting
// @Author caopengfei 2021/4/27 10:02
package changting

import (
	"math"
)

const p = 23333
const q = 10007
const n = p * q

func Decrypt(e, c int64) int64 {
	d := calD(e)
	return modPow(c, d, n)
}

func modPow(c, d, n int64) int64 {
	if n == 1 {
		return 0
	}
	r := int64(1)
	c = c % n
	for d > 0 {
		if d%2 == 1 {
			r = (r * c) % n
		}
		d = d >> 1
		c = (c * c) % n
	}
	return r
}

func calD(e int64) int64 {
	k := 1
	for {
		if (phi()*int64(k)+1)%e == 0 {
			return (phi()*int64(k) + 1) / e
		}
		k++
	}
}

func phi() int64 {
	return (p - 1) * (q - 1)
}

func mod(i int64) int64 {
	return int64(math.Mod(float64(1), float64(i)))
}
