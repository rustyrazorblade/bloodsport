package basetype

import (
	"strconv"
)

type RedisType interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
}

// base k/v structure
type BaseType struct {
	value string
}

func (b *BaseType) ToInt() (int64, error) {
	result, err := strconv.ParseInt(b.value, 10, 64)
	return result, err
}

func (b *BaseType) ToString() string {
	return string(b.value)
}

func (b *BaseType) ToFloat() (float64, error) {
	return strconv.ParseFloat(b.value, 64)
}

func (b *BaseType) IncrBy(increment int64) (string, error) {
	tmp, err := b.ToInt()

	if err != nil {
		return "", err
	}

	b.value = strconv.FormatInt(tmp + increment, 10)

	return b.value, nil
}

func (b *BaseType) DecrBy(decrement int64) (string, error) {
	tmp, err := b.ToInt()
	if err != nil {
		return "", err
	}
	b.value = strconv.FormatInt(tmp - decrement, 10)
	return b.value, nil

}

func (b *BaseType) MarshalBinary() ([]byte, error) {
	return []byte(b.value), nil
}

func (b *BaseType) UnmarshalBinary(data []byte) error {
	b.value = string(data)
	return nil
}

func NewString(value string) BaseType {
	str := BaseType{value:value}
	return str
}

