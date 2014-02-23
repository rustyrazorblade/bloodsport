package rtypes

import (
	"testing"
	. "launchpad.net/gocheck"
)

func HashTest(t *testing.T) { TestingT(t) }

type HashSuite struct {

}

func (s *HashSuite) TestHashSetAndGet(c *C) {
	h := NewHash("somekey")
	h.HSet("k", NewInteger(1, 1))
	h.HGet("k")
}

func (s *HashSuite) TestIncrBy(c *C) {
}
