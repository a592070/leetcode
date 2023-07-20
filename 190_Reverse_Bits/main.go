package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Reverse bits of a given 32 bits unsigned integer.

Note:
Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as a signed integer type. They should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above, the input represents the signed integer -3 and the output represents the signed integer -1073741825.

Input: n = 00000010100101000001111010011100
Output:    964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.

Input: n = 11111111111111111111111111111101
Output:   3221225471 (10111111111111111111111111111111)
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.

Follow up: If this function is called many times, how would you optimize it?
*/
func main() {
	//fmt.Println(reverseBits(11))       // 1011 -> 1101=13
	fmt.Println(reverseBits(43261596)) // 1011 -> 1101=13
}

func reverseBits(num uint32) uint32 {
	s := strconv.FormatUint(uint64(num), 2)

	size := len(s)
	if size < 32 {
		s = strings.Repeat("0", (32-size)) + s
	}
	size = 32

	var rs uint32
	for i := 0; i < size; i++ {
		val := s[i] - 48
		rs += uint32(math.Pow(float64(2), float64(i))) * uint32(val)
	}

	return rs
}

func reverseBits2(num uint32) uint32 {
	var rs uint32
	for i := 0; i < 32; i++ {
		rs = rs << 1
		rs = rs | (num & 1)
		num = num >> 1
	}

	return rs
}
