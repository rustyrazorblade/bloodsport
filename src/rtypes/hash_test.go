package rtypes

import (
	"testing"
	. "launchpad.net/gocheck"
)

func HashTest(t *testing.T) { TestingT(t) }

type HashSuite struct {

}

var _ = Suite(&HashSuite{})

func (s *HashSuite) TestHashSetAndGet(c *C) {
	h := NewHash("somekey")
	h.HSet("k", NewInteger(1, 1))
	result := h.HGet("k")
	num, _ := result.ToInt()
	c.Check(num, Equals, int64(1))
}

func (s *HashSuite) TestIncrBy(c *C) {
}
