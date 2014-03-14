package storage

import (
	"rtypes"
)

/*
high level database storage.  wraps the commit log and multiple files

 */

type DB struct {

}

func (db *DB) Append(key string, value rtypes.VarType) {

}

func (db *DB) Auth(password string) {

}

// noop
func (db *DB) BGRewriteAOF() {
}

// noop
func (db *DB) BGSave() {
}

func (db *DB) BitCount(key string) {

}

// returns a new string.
func (db *DB) itOp(operation string) {
}

func (db *DB) LPop(key string) {

}

// Remove and get the last element in a list, or block until one is available
func (db *DB) BRPop(key string) {

}

// Pop a value from a list, push it to another list and return it; or block until one is available
func (db *DB) BRPopLPush() {
}


// delete a key.  seems straightforward.
func (db *DB) Del(key string) {

}


func (db *DB) Incr(key string) {

}

func (db *DB) Set(key string, value rtypes.VarType, ex_seconds int, ex_milliseconds int, nx_or_xx string) {

}


