package gomisc

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHaveSameSign(c *C) {
	c.Assert(HaveSameSign(0, 1), Equals, true)
	c.Assert(HaveSameSign(0, -1), Equals, false)

	c.Assert(HaveSameSign(1, 2), Equals, true)
	c.Assert(HaveSameSign(1, -2), Equals, false)
}

func (s *MySuite) TestNoBranchAbs(c *C) {
	c.Assert(NoBranchAbs(0), Equals, int64(0))
	c.Assert(NoBranchAbs(1), Equals, int64(1))
	c.Assert(NoBranchAbs(-1), Equals, int64(1))
}

func (s *MySuite) TestIsPowOf2(c *C) {
	c.Assert(IsPowOf2(0), Equals, false)
	c.Assert(IsPowOf2(1), Equals, true)
	c.Assert(IsPowOf2(2), Equals, true)
	c.Assert(IsPowOf2(3), Equals, false)
}

func (s *MySuite) TestCountBits(c *C) {
	c.Assert(CountBits(0), Equals, 0)
	c.Assert(CountBits(1), Equals, 1)
	c.Assert(CountBits(2), Equals, 1)
	c.Assert(CountBits(3), Equals, 2)
	c.Assert(CountBits(4), Equals, 1)
}

func (s *MySuite) TestCountBitsAlt(c *C) {
	c.Assert(CountBitsAlt(0), Equals, 0)
	c.Assert(CountBitsAlt(1), Equals, 1)
	c.Assert(CountBitsAlt(2), Equals, 1)
	c.Assert(CountBitsAlt(3), Equals, 2)
	c.Assert(CountBitsAlt(4), Equals, 1)
}

func (s *MySuite) TestHasZeroByte(c *C) {
	c.Assert(HasZeroByte(0x0102030405060708), Equals, false)
	c.Assert(HasZeroByte(0x0001020304050607), Equals, true)
}

func (s *MySuite) TestHasByte(c *C) {
	c.Assert(HasByte(0x0102030405060708, 0x01), Equals, true)
	c.Assert(HasByte(0x0102030405060708, 0x09), Equals, false)
}

func (s *MySuite) TestGetLeastSignificantOneBit(c *C) {
	c.Assert(GetLeastSignificantOneBit(0x0102030405060708), Equals, uint64(0x08))
}

func (s *MySuite) TestComputeNextBitPermutation(c *C) {
	c.Assert(ComputeNextBitPermutation(1), Equals, uint64(2))
	c.Assert(ComputeNextBitPermutation(2), Equals, uint64(4))
	c.Assert(ComputeNextBitPermutation(4), Equals, uint64(8))

	c.Assert(ComputeNextBitPermutation(3), Equals, uint64(5))
	c.Assert(ComputeNextBitPermutation(5), Equals, uint64(6))
	c.Assert(ComputeNextBitPermutation(6), Equals, uint64(9))
}
