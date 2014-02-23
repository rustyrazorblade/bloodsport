package rtypes

type RedisType interface {
	Serialize(bytes []byte) *[]byte
}

type VarType interface {

}

type BaseType struct {
	timestamp int64
}

type StringType struct {
	BaseType
	value string
}

type IntType struct {
	BaseType
	value int64

}

type FloatType struct {
	BaseType
	value float64
}

type Tombstone struct {
	BaseType

}

func NewString(value string, timestamp int64) *StringType {
	str := StringType{BaseType{timestamp:timestamp}, value}
	return &str
}

func NewFloat(value float64, timestamp int64) *FloatType {
	fl := FloatType{BaseType{timestamp:timestamp}, value}
	return &fl

}

func NewInteger(value int64, timestamp int64) *IntType {
	i := IntType{BaseType{timestamp:timestamp}, value}
	return &i
}
