package hash

import (
	"rtypes/basetype"
)

type Hash struct {
	values map[string]*basetype.BaseType
	size int
}

func (hash *Hash) HDel(field string) {
	delete(hash.values, field)
}

// returns true if the key exists in the hash
func (hash *Hash) HExists(field string) bool {
	_, ok := hash.values[field]
	return ok
}

// returns the value, if it's available.  else returns... nothing?
func (hash *Hash) HGet(field string) string {
	tmp := hash.values[field]
	return tmp.ToString()
}


// increments value in a hash.  attempts to coerce the value to an int
func (hash *Hash) HIncrBy(field string, increment int64) (string, error) {
	if _, ok := hash.values[field]; !ok {
		tmp := basetype.NewString("0")
		hash.values[field] = &tmp
	}
	return hash.values[field].IncrBy(increment)
}


func (hash *Hash) HIncrByFloat(field string, increment float64) {
	if _, ok := hash.values[field]; !ok {
		tmp := basetype.NewString("0")
		hash.values[field] = &tmp
	}

}

func (hash *Hash) HKeys() {

}

func (hash *Hash) HLen() int {
	return hash.size
}

func (hash *Hash) HSet(field string, value string) {
	if _, ok := hash.values[field]; !ok {
		hash.size += 1
	}
	tmp := basetype.NewString(value)
	hash.values[field] = &tmp
}

func (hash *Hash) HSetNX(field string, value string) {

}

func (hash *Hash) HVals() {

}

func (hash *Hash) HScan(pattern string, count int) {

}

// returns a new hash, owned by a given key
// not sure if we need to save the key on here
func NewHash() *Hash {
	h := Hash{}
	h.values = make(map[string]*basetype.BaseType)
	h.size = 0

	return &h
}

