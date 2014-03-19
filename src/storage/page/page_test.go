package page


import (
	"testing"
	. "launchpad.net/gocheck"
	rstring "rtypes/string"
)

func Test(t *testing.T) { TestingT(t) }


type BaseSuite struct {}

var _ = Suite(&BaseSuite{})

func (bs *BaseSuite) TestNewPage(c *C) {
	p := NewPage()
	c.Check(p.dirty, Equals, false)
}

func (bs *BaseSuite) TestMarshal(c *C) {

}

func (bs *BaseSuite) TestUnmarshal(c *C) {

}

func (bs *BaseSuite) TestExists(c *C) {
	p := NewPage()
	p.NewString("bacon", "eggs")

	exists := p.Exists("bacon")
	c.Check(exists, Equals, true)

	b, e := p.Get("bacon")
	c.Check(e, Equals, nil)
	c.Check(b, Not(Equals), nil)

	if _, ok := b.(*rstring.String); !ok {
		// makes sure we get a string type back
		c.Fail()
	}

}
