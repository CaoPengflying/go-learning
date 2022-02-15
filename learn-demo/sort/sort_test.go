package main

import (
	"sort"
	"testing"
)

func TestBasicSort(t *testing.T) {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	t.Log(strs)

	ints := []int{5, 1, 3}
	sort.Ints(ints)
	t.Log(ints)

	s := sort.IntsAreSorted(ints)
	t.Log("sorted:", s)
}

type Student struct {
	Age int
	Name string
	Balance float32
}

func TestStructSort(t *testing.T) {
	students := []Student{{20,"cpf",100},{18,"zzc",700},{25,"cp",300}}

	sort.Slice(students, func(i, j int) bool {
		return students[i].Age < students[j].Age
	})
	sort.Slice(students, func(i, j int) bool {
		return students[i].Balance < students[j].Balance
	})
	t.Log(students)
}


