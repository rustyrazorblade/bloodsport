package string

import (
	"rtypes/basetype"
)

type String struct {
	value basetype.VarType

}

func NewString(s basetype.VarType) *String {
	new_string := String{value:s}
	return &new_string
}

func (s *String) Set(str string) {
	s.value = basetype.NewString(str)
}
