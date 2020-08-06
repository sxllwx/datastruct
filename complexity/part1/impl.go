package part1

import "fmt"

// 这个函数的时间复杂度为
// 2n*UT + 2*UT = (2n+2) *UT
func sum(n int) int {

	// 假设CPU的执行(read var, calculate, write var)时间为unit-time(UT)

	//******//
	var (
		ret = 0
		i   = 1
	)
	//******// 这两行消耗 2*UT

	//******//
	for ; i < n; i++ { // 执行 n 次
		ret += i // 执行 n 次
	}
	//******// 这两行消耗 2n*UT
	return ret
}

// 这个函数的时间复杂度为
// 2n^2 + 2n + 3
func cal(n int) int {

	// 假设CPU的执行(read var, calculate, write var)时间为unit-time(UT)

	//******//
	var (
		ret = 0
		i   = 1
		j   = 1
	)
	//******// 这两行消耗 3*UT

	//******//
	for ; i < n; i++ { // 执行 n 次
		j = 1              // 执行 n 次
		for ; j < n; j++ { // 执行 n^2 次
			ret = i * j // 执行 n^2 次
		}
	}
	//******// 这个大循环消耗，2n^2 * 2n * UT
	return ret
}

func demo(n int) {

	//******//
	var (
		sum1 = 0
		p    = 1
	)
	//******// 消耗 2*UT

	//******//
	for ; p < 100; p++ {
		sum1 += p
	}
	//******// 消耗 100*UT

	//******//
	var (
		sum2 = 0
	)
	p = 1
	//******// 消耗 2*UT

	//******//
	for ; p < n; p++ {
		sum2 += p
	}
	//******// 消耗 2n*UT

	//******//
	var (
		sum3 = 0
		i    = 1
		j    = 1
	)
	//******// 消耗 3*UT

	//******//
	for ; i < n; i++ { // 消耗 n*UT
		j = 1              // 消耗 n*UT
		for ; j < n; j++ { // 消耗 n^2*UT
			sum3 += i * j // 消耗 n^2*UT
		}
	}
	//******// 消耗 2n^2*UT + 2n*UT
}

func demo1(n int) {

	innerFunc := func(n int) int {
		var (
			sum = 0
			i   = 1
		)
		for ; i < n; i++ {
			sum += i
		}

		return sum
	}

	var (
		sum = 0
		i   = 1
	)

	for ; i < n; i++ {
		sum += innerFunc(i)

	}
}

// Summery
func Summery() {

	sum := `

	总结:
	T(n) = O(f(n))
	此处的O代表和最后的消耗时间和f(n)中的n成正比
	比如cal函数
	T(n) = O(2n^2 + 2n + 3)
	当n值到达一定的量级，T(n)的结果更大程度的取决于 n^2 
	所以最后记为T(n) = O(n^2)

	计算时间复杂度几个方便的方法:
    1. 只关注循环次数最多的代码段
	2. 总复杂度等于量级最大的那段代码的复杂度 demo 的复杂度为O(n^2)
	3. 嵌套代码的复杂度等于嵌套内外代码复杂度的乘积 demo1的时间复杂度算法为 O(n)*O(n) = O(n^2)


	2^n 和 n! 为NP问题 // 执行时间无法估量

	logN函数的时间复杂度为 O(logn)
	maddn函数的时间复杂度为o(m+n)
	`
	fmt.Println(sum)
}

func logN(n int) {

	for i := 0; i < n; i = i << 1 {
	}
}

func maddn(n, m int) int {

	var (
		sum1, sum2 = 0, 0
	)

	for i := 0; i < n; i++ {
		sum1 += i
	}
	for i := 0; i < m; i++ {
		sum2 += i
	}
	return sum1 + sum2
}
