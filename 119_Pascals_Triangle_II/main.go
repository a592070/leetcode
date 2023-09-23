package main

func main() {
}

/*
*
C(m-1)(k+1) = C(m)(k) + C(m)(k+1)
*/
func getRow(rowIndex int) []int {
	rs := make([]int, rowIndex+1)
	rs[0] = 1
	for i := 1; i <= rowIndex; i++ {
		tmp := rs[i-1]
		tmp *= rowIndex - i + 1
		tmp /= i
		rs[i] = tmp
	}
	return rs

}
func fibonacci(num int) int {
	if num <= 0 {
		return 0
	}
	if num <= 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
func power(num int) int64 {
	var rs int64 = 1
	for i := 1; i <= num; i++ {
		rs *= int64(i)
	}
	return rs
}
func compose(n, k int) int64 {
	if k == 0 {
		return 1
	}
	if k > (n / 2) {
		return compose(n, n-k)
	}

	// integer overflow
	//return power(n) / power(k) / power(n-k)

	var rs int64 = 1

	for i := 1; i <= k; i++ {
		tmp := n - i + 1
		if tmp >= 1 {
			rs *= int64(tmp)
		}
		rs /= int64(i)
	}
	return rs
}
