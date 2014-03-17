package page

/*
single page within a file
this does not manage any file pointers, or read or write directly from a file
rather, it accepts a series of bytes which are deserialized into the correct
types, and manages the memory of the page.
 */
import (
	"rtypes/basetype"
)

type Page struct {
	dirty bool
	keys map[string]*basetype.RedisType
}

func NewPage() *Page {
	page := Page{dirty:false}
	return &page
}

