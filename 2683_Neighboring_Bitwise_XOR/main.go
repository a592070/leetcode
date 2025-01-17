package main

func main() {

}

func doesValidArrayExist(derived []int) bool {
	sum := 0
	for _, v := range derived {
		sum ^= v
	}
	return sum == 0
}
