package rtypes

import (
)

type Hash struct {
	key string
	values map[string]VarType
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
func (hash *Hash) HGet(field string) VarType {
	tmp := hash.values[field]
	return tmp
}


// increments value in a hash.  attempts to coerce the value to an int
func (hash *Hash) HIncrBy(field string, increment int64) {
	if _, ok := hash.values[field];  !ok {
		hash.values[field] = NewInteger(0)
		hash.size += 1
	}
	hash.values[field].IncrBy(increment)
}


func (hash *Hash) HIncrByFloat(field string, increment float64) {

}

func (hash *Hash) HKeys() {

}

func (hash *Hash) HLen() {

}

func (hash *Hash) HSet(field string, value VarType) {
	if _, ok := hash.values[field]; !ok {
		hash.size += 1
	}
	hash.values[field] = value
}

func (hash *Hash) HSetNX(field string, value string) {

}

func (hash *Hash) HVals() {

}

func (hash *Hash) HScan() {

}

func NewHash(key string) *Hash {
	h := Hash{key:key}
	h.values = make(map[string]VarType)
	h.size = 0

	return &h
}

