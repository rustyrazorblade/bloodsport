package page


import (
	"testing"
	. "launchpad.net/gocheck"
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
	p.NewString("bacon")

}
