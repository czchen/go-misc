// https://graphics.stanford.edu/~seander/bithacks.html

package gomisc

func HaveSameSign(x, y int64) bool {
	return x^y >= 0
}

func NoBranchAbs(x int64) int64 {
	mask := x >> (64 - 1)
	return (x + mask) ^ mask
}

func IsPowOf2(x int64) bool {
	if x == 0 {
		return false
	}

	return (x & (x - 1)) == 0
}

func CountBits(x uint64) int {
	c := 0

	for ; x != 0; c++ {
		x &= x - 1
	}

	return c
}

func CountBitsAlt(x uint64) int {
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

func HasZeroByte(x uint64) bool {
	return (((x & 0x7f7f7f7f7f7f7f7f) - 0x0101010101010101) & 0x8080808080808080) != 0
}

func HasByte(x uint64, b uint8) bool {
	return HasZeroByte(x ^ (0x0101010101010101 * uint64(b)))
}

func GetLeastSignificantOneBit(x uint64) uint64 {
	return x & -x
}

func ComputeNextBitPermutation(x uint64) uint64 {
	y := GetLeastSignificantOneBit(x)
	c := x + y
	return (x^c)>>2/y | c
}
