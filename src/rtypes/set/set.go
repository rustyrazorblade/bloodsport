package set

import (

)

type Set struct {

}


// add a member to a set
func (set *Set) SAdd(member VarType) {

}

// get the number of members in a set
func (set *Set) SCard() {

}

// sdiff is a little hairy, since it may require acquiring
// the set from another node in prod
// subtracts set2 from set
func (set *Set) SDiff(set2 *Set) {

}


