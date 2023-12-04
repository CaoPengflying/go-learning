package iterator

import "testing"

func TestNewArrayIterator(t *testing.T) {
	nums := make([]interface{}, 3)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	iterator := NewArrayIterator(nums)
	for iterator.HasNext() {
		t.Log(iterator.CurrentEnum())
		iterator.Next()
	}
}
