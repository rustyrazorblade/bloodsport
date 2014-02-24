package rtypes

import (
	"testing"
	. "launchpad.net/gocheck"
)



func Test(t *testing.T) { TestingT(t) }


type BaseSuite struct {}

var _ = Suite(&BaseSuite{})


func (s *BaseSuite) TestCreate(c *C) {

	str := NewString("string", 1)
	c.Check(str.timestamp, Equals, int64(1))

	i := NewInteger(1, 1)
	c.Check(i.ToInt(), Equals, int64(1))

	f := NewFloat(1.0, 1)
	c.Check(f.ToFloat(), Equals, 1.0)
}


type IncrDecrSuite struct {
	i *IntType
	s *StringType
	f *FloatType
}

var _ = Suite(&IncrDecrSuite{})

func (s *IncrDecrSuite) SetUpTest(c *C) {
	s.i = NewInteger(10, 1)
	s.f = NewFloat(10, 1)
	s.s = NewString("hello", 1)
}

func (s *IncrDecrSuite) TestIncr(c *C) {
	i, e := s.i.IncrBy(1, 2)

	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(11))

}
func (s *IncrDecrSuite) TestDecr(c *C) {
	i, e := s.i.DecrBy(1, 2)
	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(9))

}

