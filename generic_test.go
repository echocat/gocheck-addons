package addons

import (
	"errors"
	. "gopkg.in/check.v1"
)

type CheckersTest struct{}

func init() {
	Suite(&CheckersTest{})
}

type enclosesString struct {
	string
}

func (instance enclosesString) String() string {
	return instance.string
}

func asString(what string) enclosesString {
	return enclosesString{string: what}
}

func (s *CheckersTest) TestThrowsPanicThatMatches(c *C) {
	c.Assert(func() {
		panic(errors.New("foo123"))
	}, ThrowsPanicThatMatches, "foo1.3")
	c.Assert(func() {
		panic(asString("foo123"))
	}, ThrowsPanicThatMatches, "foo1.3")
	c.Assert(func() {
		panic("foo123")
	}, ThrowsPanicThatMatches, "foo1.3")
}

func (s *CheckersTest) TestIsEmpty(c *C) {
	c.Assert("abc", Not(IsEmpty))
	c.Assert("", IsEmpty)
	c.Assert([]string{"abc"}, Not(IsEmpty))
	c.Assert([]string{}, IsEmpty)
	c.Assert([]int{1}, Not(IsEmpty))
	c.Assert([]int{}, IsEmpty)
	c.Assert(map[string]int{"abc": 1}, Not(IsEmpty))
	c.Assert(map[string]int{}, IsEmpty)
}

