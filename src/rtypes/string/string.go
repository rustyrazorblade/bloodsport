/*
The string library handles the most basic redis type

Notable commands are set, setnx, get

Unless otherwise noted, by convention, this library will accept and return
the built in string type
 */
package string

import (
	"rtypes/basetype"
	"strings"
	"encoding/gob"
)


type String struct {
	basetype.RedisDataStructureBase
	value basetype.BaseType

}

// make sure we conform to the RedisType interface
var _ basetype.RedisDataStructureInterface = &String{}

func NewString(value string) *String {
	val := basetype.NewString(value)
	new_string := String{value:val}
	return &new_string
}

func (s *String) Set(str string) {
	s.value = basetype.NewString(str)
}

func (s *String) Get() string {
	return s.value.ToString()
}


func (s *String) Append(s2 string) (int, error) {
	tmp := strings.Join([]string{s.Get(), s2}, "")
	s.Set(tmp)
	return len(tmp), nil
}

//func (s *String) Encode(encoder gob.Encoder) {
//	s.RedisDataStructureBase.Encode(encoder)
//}
