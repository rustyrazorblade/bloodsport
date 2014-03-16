package string

type String struct {
	value VarType

}

func (s *String) Set(str string) {
	s.value = NewString(str)
}
