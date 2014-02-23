package rtypes

type Hash struct {
	key string
}

func (hash *Hash) HDel(field ...string) {

}

func (hash *Hash) HExists(field string) {

}

func (hash *Hash) HGet(field string) {

}

func (hash *Hash) HGetAll() {

}

func (hash *Hash) HIncrBy(field string, increment int64) {

}

func (hash *Hash) HIncrByFloat(field string, increment float64) {

}

func (hash *Hash) HKeys() {

}

func (hash *Hash) HLen() {

}

func (hash *Hash) HMGet(field ...string) {

}

func (hash *Hash) HMSet() {

}

func (hash *Hash) HSet(field string, value string) {

}

func (hash *Hash) HSetNX(field string, value string) {

}

func (hash *Hash) HVals() {

}

func (hash *Hash) HScan() {

}

func NewHash(key string) *Hash {
	h := Hash{key}
	return &h
}

