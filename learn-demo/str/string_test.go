package str

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringCode(t *testing.T) {
	// string 的len 返回的是字节的长度，而不是字符的长度
	var str1 string
	str1 = "中"

	t.Log(str1)
	//str1[0] = "a"
	str1 = "\xE4\xB8\xA5"
	t.Log(str1)
	t.Log(len(str1))
	t.Logf("utf8 %x", str1)
	// rune函数可以转为unicode
	u1 := []rune(str1)
	t.Log(len(u1))
	t.Logf("unicode %x", u1)

}

func TestStrFn(t *testing.T) {
	str1 := "a,b,c,d"
	parts := strings.Split(str1, ",")
	t.Log(parts)
	joins := strings.Join(parts, "-")
	t.Log(joins)
	ia := "10"
	num, err := strconv.Atoi(ia)
	if err == nil {
		 t.Log(num+10)
	}
	num1 := 10
	str2 := strconv.Itoa(num1)
	str3 := str2 + "aaa"
	t.Log(str3)
}
