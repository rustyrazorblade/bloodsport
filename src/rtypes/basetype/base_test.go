package basetype

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }


type BaseSuite struct {}

var _ = Suite(&BaseSuite{})

func (s *BaseSuite) TestToInt(c *C) {
	str := NewString("1")
	result, _ := str.ToInt()
	c.Check(result, Equals, int64(1))
}

func (s *BaseSuite) TestCreate(c *C) {
	str := NewString("string")
	c.Check(str.value, Equals, "string")
}

func (s *BaseSuite) TestExpire(c *C) {

}

func (s *BaseSuite) TestMarshaling(c *C) {
	tests := []string{"hello", "test", "string", "bacon", "eggs", "whatever this is a long ass string"}
	for _, test := range tests {
		tmp := NewString(test)
		serialized_data, _ := tmp.MarshalBinary()

		tmp2 := BaseType{}
		tmp2.UnmarshalBinary(serialized_data)

		c.Check(tmp.value, Equals, tmp2.value)
	}
}

type IncrDecrSuite struct {
	s BaseType
	s2 BaseType
	s3 BaseType
}

var _ = Suite(&IncrDecrSuite{})

func (s *IncrDecrSuite) SetUpTest(c *C) {
	s.s = NewString("hello")
	s.s2 = NewString("10")
	s.s3 = NewString("10.0")
}

func (s *IncrDecrSuite) TestIncr(c *C) {
	i, e := s.s2.IncrBy(1)

	c.Check(e, Equals, nil)
	c.Check(i, Equals, "11")

	// check float, should error
	i, e = s.s3.IncrBy(1)
	c.Check(e, Not(Equals), nil)

	// check string (that cannot be coerced to an int), should fail
	i, e = s.s.IncrBy(1)
	c.Check(e, Not(Equals), nil)
}

func (s *IncrDecrSuite) TestDecr(c *C) {
	i, e := s.s2.DecrBy(1)
	c.Check(e, Equals, nil)
	c.Check(i, Equals, "9")

	// check string (that cannot be coerced to an int), should fail
	i, e = s.s.DecrBy(1)
	c.Check(e, Not(Equals), nil)
}





