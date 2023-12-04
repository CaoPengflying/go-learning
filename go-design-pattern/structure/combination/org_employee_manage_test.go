package combination

import "testing"

func TestNewOrg(t *testing.T) {
	org := NewOrg()
	println(org.Count())
}

