package gomisc

import (
	. "gopkg.in/check.v1"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHaveSameSign0(c *C) {
	c.Assert(HaveSameSign0(0, 1), Equals, true)
	c.Assert(HaveSameSign0(0, -1), Equals, false)

	c.Assert(HaveSameSign0(1, 2), Equals, true)
	c.Assert(HaveSameSign0(1, -2), Equals, false)
}

func (s *MySuite) TestNoBranchAbs0(c *C) {
	c.Assert(NoBranchAbs0(0), Equals, int64(0))
	c.Assert(NoBranchAbs0(1), Equals, int64(1))
	c.Assert(NoBranchAbs0(-1), Equals, int64(1))
}

func (s *MySuite) TestMinimum0(c *C) {
	c.Assert(Minimum0(100, 200), Equals, int64(100))
	c.Assert(Minimum0(100, -200), Equals, int64(-200))
}

func (s *MySuite) TestMaximum0(c *C) {
	c.Assert(Maximum0(100, 200), Equals, int64(200))
	c.Assert(Maximum0(100, -200), Equals, int64(100))
}

func (s *MySuite) TestIsPowOfTwo0(c *C) {
	c.Assert(IsPowOfTwo0(0), Equals, false)
	c.Assert(IsPowOfTwo0(1), Equals, true)
	c.Assert(IsPowOfTwo0(2), Equals, true)
	c.Assert(IsPowOfTwo0(3), Equals, false)
}

func (s *MySuite) TestCountBits0(c *C) {
	c.Assert(CountBits0(0), Equals, 0)
	c.Assert(CountBits0(1), Equals, 1)
	c.Assert(CountBits0(2), Equals, 1)
	c.Assert(CountBits0(3), Equals, 2)
	c.Assert(CountBits0(4), Equals, 1)
}

func (s *MySuite) TestCountBits1(c *C) {
	c.Assert(CountBits1(0), Equals, 0)
	c.Assert(CountBits1(1), Equals, 1)
	c.Assert(CountBits1(2), Equals, 1)
	c.Assert(CountBits1(3), Equals, 2)
	c.Assert(CountBits1(4), Equals, 1)
}

func (s *MySuite) TestHasZeroByte0(c *C) {
	c.Assert(HasZeroByte0(0x0102030405060708), Equals, false)
	c.Assert(HasZeroByte0(0x0001020304050607), Equals, true)
}

func (s *MySuite) TestHasByte0(c *C) {
	c.Assert(HasByte0(0x0102030405060708, 0x01), Equals, true)
	c.Assert(HasByte0(0x0102030405060708, 0x09), Equals, false)
}

func (s *MySuite) TestConditionalSet0(c *C) {
	c.Assert(ConditionalSet0(true, 0, 1), Equals, uint64(1))
	c.Assert(ConditionalSet0(false, 1, 1), Equals, uint64(0))
}

func (s *MySuite) TestGetLeastSignificantOneBit0(c *C) {
	c.Assert(GetLeastSignificantOneBit0(0x0102030405060708), Equals, uint64(0x08))
}

func (s *MySuite) TestComputeNextBitPermutation0(c *C) {
	c.Assert(ComputeNextBitPermutation0(1), Equals, uint64(2))
	c.Assert(ComputeNextBitPermutation0(2), Equals, uint64(4))
	c.Assert(ComputeNextBitPermutation0(4), Equals, uint64(8))

	c.Assert(ComputeNextBitPermutation0(3), Equals, uint64(5))
	c.Assert(ComputeNextBitPermutation0(5), Equals, uint64(6))
	c.Assert(ComputeNextBitPermutation0(6), Equals, uint64(9))
}
