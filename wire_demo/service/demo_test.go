// @Description service
// @Author caopengfei 2021/6/8 11:54

package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	discountSubs := []string{"1", "2", "3"}
	res := strings.Replace(strings.Trim(fmt.Sprint(discountSubs), "[]"), " ", ",", -1)
	t.Log(res)
	t.Log(fmt.Sprint(discountSubs))
}

func TestGetNextWeekSunday(t *testing.T) {
	assert.Nil(t, "")
	println(GetNextWeekSunday("2021-06-06").String())
	println(GetNextWeekSunday("2021-06-07").String())
	println(GetNextWeekSunday("2021-06-08").String())
	println(GetNextWeekSunday("2021-06-09").String())
	println(GetNextWeekSunday("2021-06-10").String())
	println(GetNextWeekSunday("2021-06-11").String())
	println(GetNextWeekSunday("2021-06-12").String())
	println(GetNextWeekSunday("2021-06-13").String())
}

func GetNextWeekSunday(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	zeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if zeroTime.Weekday() == time.Sunday {
		return zeroTime.AddDate(0, 0, 8)
	}
	offset := int(time.Sunday-zeroTime.Weekday()) + 7

	return zeroTime.AddDate(0, 0, 8+offset)
}
