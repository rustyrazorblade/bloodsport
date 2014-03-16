package hash

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type HashSuite struct {

}

var _ = Suite(&HashSuite{})

func (s *HashSuite) TestHashSetAndGet(c *C) {
	h := NewHash()
	h.HSet("k", "1")
	result := h.HGet("k")
	c.Check(result, Equals, "1")
}

func (s *HashSuite) TestIncrBy(c *C) {
	h := NewHash()
	h.HSet("k", "1")
	h.HIncrBy("k", 1)

	result := h.HGet("k")
	c.Check(result, Equals, "2")
}

func (s *HashSuite) TestIncrByNoVarSet(c *C) {
	h := NewHash()
	h.HIncrBy("j", 1)
	result := h.HGet("j")
	c.Check(result, Equals, "1")
}


func (s *HashSuite) TestExists(c *C) {
	h := NewHash()
	result := h.HExists("blah")
	c.Check(result, Equals, false)

	c.Check(h.size, Equals, 0)

	h.HSet("blah", "bacon")
	result = h.HExists("blah")
	c.Check(result, Equals, true)
	c.Check(h.size, Equals, 1)
}


