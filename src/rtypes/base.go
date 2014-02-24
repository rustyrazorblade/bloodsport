package rtypes

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type RedisType interface {
	Serialize(bytes []byte) *[]byte
}

type VarType interface {
}

const (
	int = iota
	float = iota
	str = iota
)
// base k/v structure
type BaseType struct {
	timestamp int64
	value []byte
	vtype int32

}

func (b *BaseType) ToInt() int64 {
	var result int64
	buf := bytes.NewReader(b.value)
	binary.Read(buf, binary.LittleEndian, &result)
	return result
}
func (b *BaseType) ToFloat() float64 {
	var result float64
	buf := bytes.NewReader(b.value)
	binary.Read(buf, binary.LittleEndian, &result)
	return result
}

func (b *BaseType) IncrBy(increment, timestamp int64) (int64, error) {
	if b.vtype == float {
		return 0, errors.New("TypeError")
	}
	val := b.ToInt() + increment

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, val)
	b.value = buf.Bytes()

	return val, nil

}

func (b *BaseType) DecrBy(decrement, timestamp int64) (int64, error) {
	if b.vtype == float {
		return 0, errors.New("TypeError")
	}
	val := b.ToInt() - decrement

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, val)
	b.value = buf.Bytes()

	return val, nil
	return 0, nil

}

type StringType struct {
	BaseType
}

type IntType struct {
	BaseType
}

type FloatType struct {
	BaseType
}

type Tombstone struct {
	BaseType

}

func NewString(value string, timestamp int64) *StringType {
	bval := []byte(value)
	str := StringType{BaseType{timestamp:timestamp, value:bval, vtype:str}}
	return &str
}

func NewFloat(value float64, timestamp int64) *FloatType {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, value)

	fl := FloatType{BaseType{timestamp:timestamp, value:buf.Bytes(), vtype:float}}
	return &fl

}

func NewInteger(value int64, timestamp int64) *IntType {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, value)
	i := IntType{BaseType{timestamp:timestamp, value:buf.Bytes(), vtype:int}}
	return &i
}
