package week1

import "errors"

func DeleteAtIndex[T any](slice []T, index int) ([]T, error) {
	if index < 0 || index > len(slice)-1 {
		return slice, errors.New("index out of range")
	}
	return append(slice[:index], slice[index+1:]...), nil
}

func ShrinkCap[T any](slice []T) []T {
	Cap, Len := cap(slice), len(slice)
	// 参数
	factor1 := 0.625
	factor2 := 0.5
	var NewCap int
	if Cap > 2048 && (Cap/Len >= 2) { // 大于2048且比例大于2，则降低为原来的 0.625
		NewCap = int(float64(factor1) * float64(Cap))
		println("shrink to 0.625x")
	} else if Cap >= 64 && Cap <= 2048 && (Cap/Len >= 4) { //容量在64～2048内且比例大于4，则降低为原来的一半
		NewCap = int(float64(factor2) * float64(Cap))
		println("shrink to 0.5x")
	} else {
		println("no need to shrink")
		return slice
	}

	newSlice := make([]T, NewCap)
	newSlice = append(newSlice, slice...)
	println(NewCap)
	return newSlice
}
