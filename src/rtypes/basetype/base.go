package basetype

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
)

type RedisType interface {
	Serialize(bytes []byte) *[]byte
}

type VarType interface {
	ToInt() (int64, error)
	IncrBy(int64) (int64, error)
	ToString() (string)
}

const (
	INT int8 = iota
	FLOAT
	STR
)
// base k/v structure
type BaseType struct {
	value []byte
	vtype int8

}

func (b *BaseType) ToInt() (int64, error) {

	if b.vtype == STR {
		result, err := strconv.Atoi(string(b.value))
		if err == nil {
			return int64(result), nil
		} else {
			return 0, err
		}
	}

	var result int64
	buf := bytes.NewReader(b.value)
	binary.Read(buf, binary.LittleEndian, &result)
	return result, nil
}

func (b *BaseType) ToString() string {
	return string(b.value)
}

func (b *BaseType) ToFloat() float64 {
	var result float64
	buf := bytes.NewReader(b.value)
	binary.Read(buf, binary.LittleEndian, &result)
	return result
}

func (b *BaseType) IncrBy(increment int64) (int64, error) {
	if b.vtype == FLOAT {
		return 0, errors.New("TypeError")
	}
	tmp, err := b.ToInt()
	if err != nil {
		return 0, err
	}
	val := tmp + increment

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, val)
	b.value = buf.Bytes()

	if b.vtype != INT {
		b.vtype = INT
	}


	return val, nil

}

func (b *BaseType) DecrBy(decrement, timestamp int64) (int64, error) {
	if b.vtype == FLOAT {
		return 0, errors.New("TypeError")
	}
	tmp, err := b.ToInt()
	if err != nil {
		return 0, err
	}
	val := tmp - decrement

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, val)
	b.value = buf.Bytes()

	return val, nil
	return 0, nil

}

func (b *BaseType) MarshalBinary() (data []byte, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, b.vtype)
	binary.Write(buf, binary.LittleEndian, b.value)
	return buf.Bytes(), nil
}

func (b *BaseType) UnmarshalBinary(data []byte) error {
	// unpack the first 2 bytes
	buf := bytes.NewBuffer(data[0:1])
	binary.Read(buf, binary.LittleEndian, &b.vtype)

	// put the rest in the value
	b.value = data[1:len(data)]
	return nil
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

func NewString(value string) *StringType {
	bval := []byte(value)
	str := StringType{BaseType{value:bval, vtype:STR}}
	return &str
}

func NewFloat(value float64) *FloatType {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, value)

	fl := FloatType{BaseType{value:buf.Bytes(), vtype:FLOAT}}
	return &fl

}

func NewInteger(value int64) *IntType {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, value)
	i := IntType{BaseType{value:buf.Bytes(), vtype:INT}}
	return &i
}
