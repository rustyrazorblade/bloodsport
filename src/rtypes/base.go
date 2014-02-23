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

func NewString(timestamp int64) *StringType {
	str := StringType{BaseType{timestamp:timestamp}}
	return &str
}
