package storage

/*
single page within a file
this does not manage any file pointers, or read or write directly from a file
rather, it accepts a series of bytes which are deserialized into the correct
types, and manages the memory of the page.
 */


type Page struct {

}

