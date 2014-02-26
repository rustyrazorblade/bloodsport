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
	h := NewHash("somekey")
	h.HSet("k", NewInteger(1, 1))
	h.HIncrBy("k", 1)

	result := h.HGet("k")
	num, _ := result.ToInt()
	c.Check(num, Equals, int64(2))
}

func (s *HashSuite) TestIncrByNoVarSet(c *C) {
	h := NewHash("somekey")
	h.HIncrBy("j", 1)
	result := h.HGet("j")
	num, _ := result.ToInt()
	c.Check(num, Equals, int64(1))
}

