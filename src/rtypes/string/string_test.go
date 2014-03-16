package string

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type StringSuite struct {

}

var _ = Suite(&StringSuite{})

func (suite *StringSuite) TestSet(c *C) {
	s := NewString()
	s.Set("test")
	result := s.Get()
	c.Check(result, Equals, "test")
}

