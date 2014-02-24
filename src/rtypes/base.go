package rtypes

import (
	"bytes"
	"encoding/binary"
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

func (bt *BaseType) GetValue() []byte {
	return bt.value
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
