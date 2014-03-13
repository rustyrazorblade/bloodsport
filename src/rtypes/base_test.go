package rtypes

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
	c.Check(str.vtype, Equals, STR)

	i := NewInteger(1)
	result, _ := i.ToInt()
	c.Check(result, Equals, int64(1))

	f := NewFloat(1.0)
	c.Check(f.ToFloat(), Equals, 1.0)
}


func (s *BaseSuite) TestMarshaling(c *C) {
	tests := []string{"hello", "test", "string", "bacon", "eggs", "whatever this is a long ass string"}
	for _, test := range tests {
		tmp := NewString(test)
		serialized_data, err := tmp.MarshalBinary()

		c.Check(err, Equals, nil)

		tmp2 := BaseType{}
		tmp2.UnmarshalBinary(serialized_data)

		c.Check(err, Equals, nil)
		c.Check(tmp2.vtype, Equals, tmp.vtype)

		for i := range tmp.value {
			c.Check(tmp2.value[i], Equals, tmp.value[i])
		}

	}

}


type IncrDecrSuite struct {
	i *IntType
	s *StringType
	s2 *StringType
	f *FloatType
}

var _ = Suite(&IncrDecrSuite{})

func (s *IncrDecrSuite) SetUpTest(c *C) {
	s.i = NewInteger(10)
	s.f = NewFloat(10)

	s.s = NewString("hello")
	s.s2 = NewString("10")
}

func (s *IncrDecrSuite) TestIncr(c *C) {
	i, e := s.i.IncrBy(1)

	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(11))

	// check float, should error
	i, e = s.f.IncrBy(1)
	c.Check(e, Not(Equals), nil)

	// check string (that cannot be coerced to an int), should fail
	i, e = s.s.IncrBy(1)
	c.Check(e, Not(Equals), nil)

	// check string (that can be coerced to an int), should be ok
	i, e = s.s2.IncrBy(1)
	c.Check(e, Equals, nil)
	c.Check(i, Equals, int64(11))
	// field should be a string now
	c.Check(s.s2.vtype, Equals, INT)
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





