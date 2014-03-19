package page

/*
single page within a file
this does not manage any file pointers, or read or write directly from a file
rather, it accepts a series of bytes which are deserialized into the correct
types, and manages the memory of the page.
 */

import (
	"rtypes/basetype"
	rstring "rtypes/string"
)

type Page struct {
	dirty bool
	keys map[string]basetype.RedisType
}

func NewPage() *Page {
	page := Page{dirty:false}
	page.keys = make(map[string]basetype.RedisType)
	return &page
}

func (p *Page) Exists(key string) bool {
	_, present := p.keys[key]
	return present
}

func (p *Page) NewString(key string) *rstring.String {
	tmp := rstring.NewString()
	p.keys[key] = tmp
	return tmp
}
