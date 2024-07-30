package main

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

type LinkedListV1[T any] struct {
	head *nodeV1[T]
}

func (l *LinkedListV1[T]) Add(index int, val T) {}

type nodeV1[T any] struct {
	val T
}

func UseList() {
	l := &LinkedListV1[int]{}
	l.Add(1, 12)
}

func Sum[A Number](vals []A) A {
	var ans A
	for _, val := range vals {
		ans += val
	}
	return ans
}

type Integer int

type Number interface { // 声明类型约束
	~int | uint | float32 | float64 // 没有波浪线的话，衍生类型会报错
}

func UseSum() {
	res := Sum[int]([]int{1, 2})
	println(res)
	resV1 := Sum[Integer]([]Integer{1, 2})
	println(resV1)
}
