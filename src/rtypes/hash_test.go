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
	h.HSet("k", NewInteger(1))
	result := h.HGet("k")
	num, _ := result.ToInt()
	c.Check(num, Equals, int64(1))
}

func (s *HashSuite) TestIncrBy(c *C) {
	h := NewHash("somekey")
	h.HSet("k", NewInteger(1))
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


func (s *HashSuite) TestExists(c *C) {
	h := NewHash("test")
	result := h.HExists("blah")
	c.Check(result, Equals, false)

	c.Check(h.size, Equals, 0)

	h.HSet("blah", NewString("bacon"))
	result = h.HExists("blah")
	c.Check(result, Equals, true)
	c.Check(h.size, Equals, 1)
}


