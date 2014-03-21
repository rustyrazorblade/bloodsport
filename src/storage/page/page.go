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
	keys map[string]basetype.RedisDataStructureInterface
}

func NewPage() *Page {
	page := Page{dirty:false}
	page.keys = make(map[string]basetype.RedisDataStructureInterface)
	return &page
}

type Command func()

func (p *Page) Exists(key string) bool {
	_, present := p.keys[key]
	return present
}

func (p *Page) NewString(key string, value string) *rstring.String {
	tmp := rstring.NewString(value)
	p.keys[key] = tmp
	return tmp
}

func (p *Page) Get(key string) (basetype.RedisDataStructureInterface, error) {
	// returns an interface.  it's up to the higher layer to decide what to do with it for now
	// if the key is marked with an expiration that's before right now, we expire the key,
	// delete it, and return nil
	if tmp, ok := p.keys[key]; ok {
		if tmp.Expired() {
			// todo check for memory leaks
			delete(p.keys, key)
			return nil, nil
		}
		return tmp, nil
	} else  {
		return nil, nil
	}
	return nil, nil

}


func (p *Page) Execute(c Command) {
	c()
}

func (p *Page) Start() {
	// starts and actor
}

