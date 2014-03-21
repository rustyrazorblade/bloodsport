package page


import (
	"testing"
	. "launchpad.net/gocheck"
	rstring "rtypes/string"
	"time"
	"fmt"
)

func Test(t *testing.T) { TestingT(t) }


type BaseSuite struct {}

var _ = Suite(&BaseSuite{})

func (bs *BaseSuite) TestNewPage(c *C) {
	p := NewPage()
	c.Check(p.dirty, Equals, false)
}

func (bs *BaseSuite) TestMarshaling(c *C) {

}


func (bs *BaseSuite) TestExists(c *C) {
	p := NewPage()
	p.NewString("bacon", "eggs")

	exists := p.Exists("bacon")
	c.Check(exists, Equals, true)

	b, e := p.Get("bacon")
	c.Check(e, Equals, nil)
	c.Check(b, Not(Equals), nil)

	if b_as_string, ok := b.(*rstring.String); ok {
		// makes sure we get a string type back
		value := b_as_string.Get()
		c.Check(value, Equals, "eggs")

	} else {
		c.Fail()
	}

}


func (bs *BaseSuite) TestExpireInFuture(c *C) {
	the_future := time.Now().Add(time.Duration(10) * time.Second)

	p := NewPage()
	s := p.NewString("frank_dux", "wins")
	s.SetExpire(&the_future)
	s2, _ := p.Get("frank_dux")
	c.Check(s2, Not(Equals), nil)
}

func (bs *BaseSuite) TestExpireInPast(c *C) {
	// test the expired key case
	p := NewPage()
	the_past := time.Now().Add(-time.Duration(10) * time.Second)
	s := p.NewString("chong_li", "loses")
	s.SetExpire(&the_past)
	s2, _ := p.Get("chong_li")
	c.Check(s2, Equals, nil)
}

func (bs *BaseSuite) TestExecute(c *C) {
	command := func () {
		fmt.Println("OK")
	}
	p := NewPage()
	p.Execute((Command)(command))
}
