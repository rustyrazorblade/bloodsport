package basetype

import (
	"strconv"
)

type RedisType interface {
	Serialize(bytes []byte) *[]byte
}

// base k/v structure
type BaseType struct {
	value string
}

func (b *BaseType) ToInt() (int, error) {
	result, err := strconv.Atoi(b.value)
	return result, err
}

func (b *BaseType) ToString() string {
	return string(b.value)
}

func (b *BaseType) ToFloat() (float64, error) {
	return strconv.ParseFloat(b.value, 64)
}

func (b *BaseType) IncrBy(increment int) (string, error) {
	tmp, err := b.ToInt()

	if err != nil {
		return "", err
	}
	b.value = strconv.Itoa(tmp + increment)

	return b.value, nil
}

func (b *BaseType) DecrBy(decrement int) (string, error) {
	tmp, err := b.ToInt()
	if err != nil {
		return "", err
	}
	b.value = strconv.Itoa(tmp - decrement)
	return b.value, nil

}

func (b *BaseType) MarshalBinary() (data []byte) {
	return []byte(b.value)
}

func (b *BaseType) UnmarshalBinary(data []byte) error {
	b.value = string(data)
	return nil
}

func NewString(value string) BaseType {
	str := BaseType{value:value}
	return str
}

