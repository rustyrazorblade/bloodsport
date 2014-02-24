package rtypes

import (
	"fmt"
	"testing"
	. "launchpad.net/gocheck"
)

func HashTest(t *testing.T) { TestingT(t) }

type HashSuite struct {

}

func (s *HashSuite) TestHashSetAndGet(c *C) {
	h := NewHash("somekey")
	h.HSet("k", NewInteger(1, 1))
	result := h.HGet("k")
	fmt.Println(result)
}

func (s *HashSuite) TestIncrBy(c *C) {
}
