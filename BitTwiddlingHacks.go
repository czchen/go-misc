// https://graphics.stanford.edu/~seander/bithacks.html

package gomisc

func HaveSameSign0(x, y int64) bool {
	return x^y >= 0
}

func NoBranchAbs0(x int64) int64 {
	mask := x >> (64 - 1)
	return (x + mask) ^ mask
}

func Minimum0(x, y int64) int64 {
	var flag int64

	if x < y {
		flag = 1
	} else {
		flag = 0
	}

	return y ^ ((x ^ y) & -flag)
}

func Maximum0(x, y int64) int64 {
	var flag int64

	if x < y {
		flag = 1
	} else {
		flag = 0
	}

	return x ^ ((x ^ y) & -flag)
}


func IsPowOfTwo0(x int64) bool {
	if x == 0 {
		return false
	}

	return (x & (x - 1)) == 0
}

func CountBits0(x uint64) int {
	c := 0

	for ; x != 0; c++ {
		x &= x - 1
	}

	return c
}

func CountBits1(x uint64) int {
	S := [...]uint{1, 2, 4, 8, 16, 32}
	B := [...]uint64{
		0x5555555555555555,
		0x3333333333333333,
		0x0f0f0f0f0f0f0f0f,
		0x00ff00ff00ff00ff,
		0x0000ffff0000ffff,
		0x00000000ffffffff,
	}

	for i := 0; i < len(S); i++ {
		x = ((x >> S[i]) & B[i]) + (x & B[i])
	}

	return int(x)
}

func HasZeroByte0(x uint64) bool {
	return (((x & 0x7f7f7f7f7f7f7f7f) - 0x0101010101010101) & 0x8080808080808080) != 0
}

func HasByte0(x uint64, b uint8) bool {
	return HasZeroByte0(x ^ (0x0101010101010101 * uint64(b)))
}

func ConditionalSet0(cond bool, mask, value uint64) uint64 {
	var flag uint64

	if cond {
		flag = 1
	} else {
		flag = 0
	}

	value ^= (-flag ^ value) & mask;

	return value
}

func GetLeastSignificantOneBit0(x uint64) uint64 {
	return x & -x
}

func ComputeNextBitPermutation0(x uint64) uint64 {
	y := GetLeastSignificantOneBit0(x)
	c := x + y
	return (x^c)>>2/y | c
}
