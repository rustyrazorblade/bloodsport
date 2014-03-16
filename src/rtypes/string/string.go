/*
The string library handles the most basic redis type

Notable commands are set, setnx, get

Unless otherwise noted, by convention, this library will accept and return
the built in string type
 */
package string

import (
	"rtypes/basetype"
)

type String struct {
	value basetype.VarType

}

func NewString() *String {
	new_string := String{}
	return &new_string
}

func (s *String) Set(str string) {
	s.value = basetype.NewString(str)
}

func (s *String) Get() string {
	return s.value.ToString()
}
