package rtypes

import (
	"testing"
	. "launchpad.net/gocheck"
)



func Test(t *testing.T) { TestingT(t) }


type BaseSuite struct {}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) TestToInt(c *C) {
	str := NewString("1", 1)
	result, _ := str.ToInt()
	c.Check(result, Equals, int64(1))


}

func (s *BaseSuite) TestCreate(c *C) {

	str := NewString("string", 1)
	c.Check(str.timestamp, Equals, int64(1))

	i := NewInteger(1, 1)
	result, _ := i.ToInt()
	c.Check(result, Equals, int64(1))

	f := NewFloat(1.0, 1)
	c.Check(f.ToFloat(), Equals, 1.0)
}


type IncrDecrSuite struct {
	i *IntType
	s *StringType
	s2 *StringType
	f *FloatType
}

var _ = Suite(&IncrDecrSuite{})

func (s *IncrDecrSuite) SetUpTest(c *C) {
	s.i = NewInteger(10, 1)
	s.f = NewFloat(10, 1)

	s.s = NewString("hello", 1)
	s.s2 = NewString("10", 1)
}

func (s *IncrDecrSuite) TestIncr(c *C) {
	i, e := s.i.IncrBy(1, 2)

	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(11))

	// check float, should error
	i, e = s.f.IncrBy(1, 1)
	c.Check(e, Not(Equals), nil)

	// check string (that cannot be coerced to an int), should fail
	i, e = s.s.IncrBy(1, 1)
	c.Check(e, Not(Equals), nil)

	// check string (that can be coerced to an int), should be ok
	i, e = s.s2.IncrBy(1, 1)
	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(11))
}

func (s *IncrDecrSuite) TestDecr(c *C) {
	i, e := s.i.DecrBy(1, 2)
	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(9))

	// check string (that cannot be coerced to an int), should fail
	i, e = s.s.DecrBy(1, 1)
	c.Check(e, Not(Equals), nil)

	// check string (that can be coerced to an int), should be ok
	i, e = s.s2.DecrBy(1, 1)
	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(9))
}


