package singleton

import "testing"

func TestGetSingleton(t *testing.T) {
	session := GetSingleton()
	s2 := GetSingleton()
	t.Log(session)
	t.Log(s2)
	t.Log(s2 == session)
}

