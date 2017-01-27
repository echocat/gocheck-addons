package addons

import (
	. "gopkg.in/check.v1"
)

type StringsTest struct{}

func init() {
	Suite(&StringsTest{})
}

func (s *StringsTest) TestContains(c *C) {
	c.Assert("abc", Contains, "a")
	c.Assert("abc", Contains, "b")
	c.Assert("abc", Contains, "c")
	c.Assert("abc", Contains, "abc")
	c.Assert(asString("abc"), Contains, asString("b"))
	c.Assert("abc", Contains, "")
	c.Assert("abc", Not(Contains), "x")
}

func (s *StringsTest) TestHasPrefix(c *C) {
	c.Assert("abc", HasPrefix, "a")
	c.Assert("abc", HasPrefix, "ab")
	c.Assert("abc", HasPrefix, "abc")
	c.Assert(asString("abc"), HasPrefix, asString("a"))
	c.Assert("abc", HasPrefix, "")
	c.Assert("abc", Not(HasPrefix), "b")
	c.Assert("abc", Not(HasPrefix), "c")
}

func (s *StringsTest) TestHasSuffix(c *C) {
	c.Assert("abc", HasSuffix, "c")
	c.Assert("abc", HasSuffix, "bc")
	c.Assert("abc", HasSuffix, "abc")
	c.Assert(asString("abc"), HasSuffix, asString("c"))
	c.Assert("abc", HasSuffix, "")
	c.Assert("abc", Not(HasSuffix), "a")
	c.Assert("abc", Not(HasSuffix), "b")
}
