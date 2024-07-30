package week1

import "testing"

func TestDelete(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	newslice, err := DeleteAtIndex[int](slice, 3)
	if err != nil {
		t.Error(err)
	}
	t.Log(newslice)

	s1 := make([]int, 10, 20)
	s2 := make([]int, 1000, 5000)
	s3 := make([]int, 400, 2000)
	t.Log(cap(s1), cap(s2), cap(s3))

	s1 = ShrinkCap[int](s1)
	s2 = ShrinkCap[int](s2)
	println(cap(s2))
	s3 = ShrinkCap[int](s3)

	t.Log(cap(s1), cap(s2), cap(s3))

}
