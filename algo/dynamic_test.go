// @Description algo
// @Author caopengfei 2021/8/12 10:49

package algo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//1. 确定最后一步，假设使用x个操作达到最后一步
//2. 拆成子问题，降阶。fn 与 fn-1 之间的关系
//3. 初始条件和边界
//4. 计算顺序

//跳跃次数

func TestJumpCount(t *testing.T) {
	println(jumpCount([]int{2, 3, 1}))
}

func jumpCount(nums []int) int {
	index := 0
	step := 0
	for index < len(nums)-1 {
		if nums[index]+index >= len(nums)-1 {
			step++
			return step
		}
		length := nums[index]
		maxIndex := 0
		tempIndex := 0
		for i := 0; i < length; i++ {
			if index+i+1 >= len(nums) {
				step++
				return step
			}
			if maxIndex < nums[index+i+1]+index+i+1 {
				maxIndex = nums[index+i+1] + index + i + 1
				tempIndex = index + i + 1
			}
		}
		index = tempIndex
		step++
	}
	return step
}

//最长回文子串
func TestLongSubStr(t *testing.T) {
	println(longSubStr("bbddc"))
}

func longSubStr(str string) string {
	if len(str) < 2 {
		return str
	}

	dp := make([][]bool, len(str))
	for i, _ := range dp {
		dp[i] = make([]bool, len(str))
		dp[i][i] = true
	}

	maxLen := 0
	begin := 0

	for l := 2; l < len(str); l++ { //子串的长度
		for left := 0; left < len(str); left++ { //子串的左下标
			right := left + l - 1
			if right >= len(str) {
				break
			}

			if str[left] != str[right] {
				dp[left][right] = false
			} else {
				if right-left < 3 {
					dp[left][right] = true
				} else {
					dp[left][right] = dp[left+1][right-1]
				}
			}

			if l > maxLen && dp[left][right] {
				maxLen = l
				begin = left
			}

		}
	}

	return str[begin : begin+maxLen]
}

// 有n块石头分别在x轴的0,1,...,n-1位置
//一只青蛙在石头0,想跳到石头n-1
//如果青蛙在第i块石头上，它最多可以向右跳距离ai
//问青蛙能否跳到石头n-1

func TestJump(t *testing.T) {
	println(jump([]int{3, 2, 1, 0, 4}))
}

func jump(stones []int) bool {
	arr := make([]bool, len(stones))
	arr[0] = true
	for i, stone := range stones {

		for j := 1; j <= stone; j++ {
			if i+j < len(stones) {
				arr[i+j] = true
			}
		}
	}
	return arr[len(stones)-1]
}

// 有一个i,j的方格，一个机器人只能往右走或者往下下，一共有多少种方式能够到达右下角
func TestRobot(t *testing.T) {
	println(robot(2, 3))
}

func robot(x, y int) int {
	grid := make([][]int, x)
	for index, _ := range grid {
		grid[index] = make([]int, y)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
				continue
			}
			grid[i][j] = grid[i-1][j] + grid[i][j-1]
		}
	}
	return grid[x-1][y-1]
}

// 有n枚硬币 2，5，7 请问最少能用多少个硬币来找开 27 m

func TestCoinChange(t *testing.T) {
	println(coinChange([]int{2, 5, 7}, 27))
}

func coinChange(coins []int, money int) int {
	arr := make([]int, money+1)

	arr[0] = 0

	for i := 1; i <= money; i++ {
		arr[i] = math.MaxInt32
		for _, coin := range coins {
			if i >= coin && arr[i-coin] != math.MaxInt32 {
				arr[i] = min(arr[i-coin]+1, arr[i])
			}
		}
	}
	if arr[money] == math.MaxInt32 {
		return -1
	}
	return arr[money]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestA(t *testing.T) {
	titleIdList := []int{48}
	for i := 47; i >= 1; i-- {
		titleIdStr := ""
		for _, titleId := range titleIdList {
			if titleIdStr != "" {
				titleIdStr += ","
			}
			titleIdStr += strconv.Itoa(titleId)
		}
		fmt.Printf("SELECT %d,count(DISTINCT stu_id) from wx_odsdb.ods_ms_xes_studycenter_read_practice_xes_read_practice_user_title_fd WHERE title_id = %d and stu_id not in (SELECT stu_id from wx_odsdb.ods_ms_xes_studycenter_read_practice_xes_read_practice_user_title_fd where title_id in (%s))", i,i, titleIdStr)
		fmt.Println()
		fmt.Println("union all")
		titleIdList = append(titleIdList, i)
	}
}
