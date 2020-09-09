package thinking

//写一个函数，输入 x，输出斐波那契数列中第 x 位的元素。
//例如，输入 4，输出 2；输入 9，输出 21。要求：需要用递归的方式来实现。

func fib(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
