// @Description changting
// @Author caopengfei 2021/4/27 09:45
package changting

func MergeNums(nums []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res = append(res, nums[i]+nums[j])
		}
	}

	return res
}

func CountDivide3(nums []int) int {
	count := 0
	for _, num := range nums {
		if num%3 == 0 {
			count++
		}
	}
	return count
}
