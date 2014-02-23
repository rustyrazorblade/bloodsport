package rtypes

import (
	"testing"
	. "launchpad.net/gocheck"
)



func BaseTest(t *testing.T) { TestingT(t) }


type BaseSuite struct {


}


func (s *BaseSuite) TestCreate(c *C) {
	st := StringType{BaseType{timestamp:1}}
	it := IntType{BaseType{timestamp:1}}
	c.Check(st, Not(Equals), it)
	c.Check(st.timestamp, Equals, 1)
	c.Check(it.timestamp, Equals, 1)

	str2 := NewString(1)
	c.Check(str2.timestamp, Equals, 1)
}

