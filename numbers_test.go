package addons

import (
	. "gopkg.in/check.v1"
	"time"
)

type NumbersTest struct{}

func init() {
	Suite(&NumbersTest{})
}

func (s *NumbersTest) TestIsLessThan(c *C) {
	c.Assert(22, IsLessThan, 66)
	c.Assert(22, Not(IsLessThan), 11)
	c.Assert(22, Not(IsLessThan), 22)
	c.Assert(22.5, IsLessThan, 66.5)
	c.Assert(22.5, Not(IsLessThan), 11.5)
	c.Assert(22.5, Not(IsLessThan), 22.5)
	c.Assert(time.Duration(22), IsLessThan, time.Duration(66))

	result, message := IsLessThan.Check([]interface{}{22.5, 11}, []string{"obtained", "compareTo"})
	c.Assert(result, Equals, false)
	c.Assert(message, Equals, "'obtained' type not equal to the type to 'compareTo' type.")
}

func (s *NumbersTest) TestIsLessThanOrEqual(c *C) {
	c.Assert(22, IsLessThanOrEqualTo, 66)
	c.Assert(22, IsLessThanOrEqualTo, 22)
	c.Assert(22, Not(IsLessThanOrEqualTo), 11)
	c.Assert(22.5, IsLessThanOrEqualTo, 66.5)
	c.Assert(22.5, IsLessThanOrEqualTo, 22.5)
	c.Assert(22.5, Not(IsLessThanOrEqualTo), 11.5)
	c.Assert(time.Duration(22), IsLessThanOrEqualTo, time.Duration(66))
	c.Assert(time.Duration(22), IsLessThanOrEqualTo, time.Duration(22))

	result, message := IsLessThanOrEqualTo.Check([]interface{}{22.5, 11}, []string{"obtained", "compareTo"})
	c.Assert(result, Equals, false)
	c.Assert(message, Equals, "'obtained' type not equal to the type to 'compareTo' type.")
}

func (s *NumbersTest) TestIsLargerThan(c *C) {
	c.Assert(22, IsLargerThan, 11)
	c.Assert(22, Not(IsLargerThan), 66)
	c.Assert(22, Not(IsLargerThan), 22)
	c.Assert(66.5, IsLargerThan, 22.5)
	c.Assert(11.5, Not(IsLargerThan), 22.5)
	c.Assert(22.5, Not(IsLargerThan), 22.5)
	c.Assert(time.Duration(22), IsLargerThan, time.Duration(11))

	result, message := IsLargerThan.Check([]interface{}{22.5, 11}, []string{"obtained", "compareTo"})
	c.Assert(result, Equals, false)
	c.Assert(message, Equals, "'obtained' type not equal to the type to 'compareTo' type.")
}

func (s *NumbersTest) TestIsLargerThanOrEqualTo(c *C) {
	c.Assert(22, IsLargerThanOrEqualTo, 11)
	c.Assert(22, Not(IsLargerThanOrEqualTo), 66)
	c.Assert(22, IsLargerThanOrEqualTo, 22)
	c.Assert(66.5, IsLargerThanOrEqualTo, 22.5)
	c.Assert(11.5, Not(IsLargerThanOrEqualTo), 22.5)
	c.Assert(22.5, IsLargerThanOrEqualTo, 22.5)
	c.Assert(time.Duration(22), IsLargerThanOrEqualTo, time.Duration(11))
	c.Assert(time.Duration(22), IsLargerThanOrEqualTo, time.Duration(22))

	result, message := IsLargerThanOrEqualTo.Check([]interface{}{22.5, 11}, []string{"obtained", "compareTo"})
	c.Assert(result, Equals, false)
	c.Assert(message, Equals, "'obtained' type not equal to the type to 'compareTo' type.")
}
