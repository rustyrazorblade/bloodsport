package string

import (
	"testing"
	. "launchpad.net/gocheck"
)

func StringTest(t *testing.T) { TestingT(t) }

type StringTestSuite struct {

}

var _ = Suite(&StringTestSuite{})
