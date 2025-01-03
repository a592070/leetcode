package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {

}

// Greedy algo.
// failed to find optimized solution. [186,419,83,408] 6249 -> 20
func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	slices.Sort(coins)

	rs := 0
	for amount > 0 {
		if amount < coins[0] {
			return -1
		}
		for i := len(coins) - 1; i > 0; i-- {
			if coins[i] <= amount {
				amount = amount - coins[i]
				rs++
				break
			}
		}
	}

	return rs
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
consider coins[1,2,5]
f(0) = 0
f(1) = 1 = min( 1+f(1-1) )
f(2) = 1 = min( 1+f(2-1), 1+f(2-2) )
f(3) = 2 = min( 1+f(3-1), 1+f(3-2) )
f(4) = 2 = min( 1+f(4-1), 1+f(4-2) )
f(5) = 1 = min( 1+f(5-1), 1+f(5-2), 1+f(5-5) )
f(6) = 2 = min( 1+f(6-1), 1+f(6-2), 1+f(6-5) )
f(11) = 3 = min( 1+f(11-1), 1+f(11-2), 1+f(11-5) )

so. for coins in [c1, c2, ... cn]
f(m) = min( 1+f(m-c1), 1+f(m-c2), ... , 1+f(m-cn) )
*/
func coinChange3(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}

	memorizedOptimization := make([]int, amount+1)
	memorizedOptimization[0] = 0

	for i := 1; i <= amount; i++ {
		memorizedOptimization[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				memorizedOptimization[i] = Min(memorizedOptimization[i], 1+memorizedOptimization[i-coin])
			}
		}
	}

	fmt.Println(memorizedOptimization)
	if memorizedOptimization[amount] == math.MaxInt32 {
		return -1
	}

	return memorizedOptimization[amount]
}
