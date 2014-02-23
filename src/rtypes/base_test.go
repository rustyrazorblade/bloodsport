package rtypes

import (
	"testing"
	. "launchpad.net/gocheck"
)



func BaseTest(t *testing.T) { TestingT(t) }


type BaseSuite struct {


}


func (s *BaseSuite) TestCreate(c *C) {

	str := NewString("string", 1)
	c.Check(str.timestamp, Equals, 1)

	i := NewInteger(1, 1)
	c.Check(i.value, Equals, 1)

	f := NewFloat(1.0, 1)
	c.Check(f.value, Equals, 1.0)
}

