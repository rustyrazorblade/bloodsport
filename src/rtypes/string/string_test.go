package string

import (
	"testing"
	. "launchpad.net/gocheck"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type StringSuite struct {

}

var _ = Suite(&StringSuite{})

func (suite *StringSuite) TestSet(c *C) {
	s := NewString("banana")
	s.Set("test")
	result := s.Get()
	c.Check(result, Equals, "test")
}


func (suite *StringSuite) TestExpire(c *C) {

	the_future := time.Now().Add(time.Duration(10) * time.Second)
	s := NewString("van_damme")
	s.SetExpire(&the_future)
}
