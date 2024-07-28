package main

func DeferClosureV1() {
	i := 0
	defer func() {
		println(i) // 闭包
	}()
	i = 1
}

func DeferClosureV2() {
	i := 0
	defer func(val int) {
		println(val)
	}(i) // 参数传递
	i = 1
}

func DeferReturnV1() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a // return先执行,int类型的返回值在defer前被确定
}

func DeferReturnV2() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a // return 先执行，返回a=0；但defer又将命名返回值a重新赋为1，然后函数关闭
}

func DeferCloseLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func DeferCloseLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println(val)
		}(i)
	}
}

func DeferCloseLoopV3() {
	for i := 0; i < 10; i++ {
		j := i //j不被闭包 使用完就不会被清除，所以一共存了0～9的j变量
		defer func() {
			println(j)
		}()
	}
}
