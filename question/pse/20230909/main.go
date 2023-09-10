package main

func main() {

}

/*
*
Is able represent money by 3, 5, 7 coins?
*/
func availableRepresentMoney(money int) bool {
	if money == 0 {
		return true
	}
	temp := make([]bool, money+1)
	temp[0] = true
	coins := []int{3, 5, 7}
	for _, coin := range coins {
		for i := coin; i <= money; i++ {
			temp[i] = temp[i-coin]
		}
	}
	return temp[money]
}
