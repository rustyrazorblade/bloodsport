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
