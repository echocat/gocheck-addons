[![Go Report Card](https://goreportcard.com/badge/github.com/echocat/gocheck-addons)](https://goreportcard.com/report/github.com/echocat/gocheck-addons)
[![Circle CI](https://img.shields.io/circleci/project/echocat/gocheck-addons.svg?style=flat-square)](https://circleci.com/gh/echocat/gocheck-addons)
[![Coverage Status](https://img.shields.io/coveralls/echocat/gocheck-addons/master.svg?style=flat-square)](https://coveralls.io/github/echocat/gocheck-addons?branch=master)
[![License](https://img.shields.io/github/license/echocat/gocheck-addons.svg?style=flat-square)](LICENSE)

# gocheck-addons

Provides some addons functions for usage together with [gocheck](http://labix.org/gocheck).

* [Installation](#installation)
* [Checkers](#checkers)
* [Run tests of this library](#run-tests-of-this-library)
* [Contributing](#contributing)
* [Support](#support)
* [License](#license)

## Installation

Install dependencies:
```bash
go get -u github.com/echocat/gocheck-addons
```

Import in source code files:
```golang
import (
	. "github.com/echocat/gocheck-addons"
	. "gopkg.in/check.v1"
)
```

## Checkers

### ThrowsPanicThatMatches

Check for a panic message that should occur.

Example:
```golang
c.Assert(func() {
    panic(errors.New("foo123"))
}, ThrowsPanicThatMatches, "foo1.3")
c.Assert(func() {
    panic(enclosesString{string: "foo123"})
}, ThrowsPanicThatMatches, "foo1.3")
c.Assert(func() {
    panic("foo123")
}, ThrowsPanicThatMatches, "foo1.3")
```

### IsEmpty

Check that a value is empty. Currently support strings, arrays and maps.

Example:
```golang
c.Assert("abc", Not(IsEmpty))
c.Assert("", IsEmpty)
c.Assert([]string{"abc"}, Not(IsEmpty))
c.Assert([]string{}, IsEmpty)
c.Assert([]int{1}, Not(IsEmpty))
c.Assert([]int{}, IsEmpty)
c.Assert(map[string]int{"abc": 1}, Not(IsEmpty))
c.Assert(map[string]int{}, IsEmpty)
```

### IsLessThan

Check that a value is less than to provided.

Example:
```golang
c.Assert(22, IsLessThan, 66)
c.Assert(22, Not(IsLessThan), 11)
c.Assert(22, Not(IsLessThan), 22)
c.Assert(22.5, IsLessThan, 66.5)
c.Assert(22.5, Not(IsLessThan), 11.5)
c.Assert(22.5, Not(IsLessThan), 22.5)
c.Assert(time.Duration(22), IsLessThan, time.Duration(66))
```

### IsLessThanOrEqual

Check that a value is less than or equal to provided.

Example:
```golang
c.Assert(22, IsLessThanOrEqualTo, 66)
c.Assert(22, IsLessThanOrEqualTo, 22)
c.Assert(22, Not(IsLessThanOrEqualTo), 11)
c.Assert(22.5, IsLessThanOrEqualTo, 66.5)
c.Assert(22.5, IsLessThanOrEqualTo, 22.5)
c.Assert(22.5, Not(IsLessThanOrEqualTo), 11.5)
c.Assert(time.Duration(22), IsLessThanOrEqualTo, time.Duration(66))
c.Assert(time.Duration(22), IsLessThanOrEqualTo, time.Duration(22))
```

### IsLargerThan

Check that a value is larger than to provided.

Example:
```golang
c.Assert(22, IsLargerThan, 11)
c.Assert(22, Not(IsLargerThan), 66)
c.Assert(22, Not(IsLargerThan), 22)
c.Assert(66.5, IsLargerThan, 22.5)
c.Assert(11.5, Not(IsLargerThan), 22.5)
c.Assert(22.5, Not(IsLargerThan), 22.5)
c.Assert(time.Duration(22), IsLargerThan, time.Duration(11))
```

### IsLargerThanOrEqualTo

Check that a value is larger than or equal to provided.

Example:
```golang
c.Assert(22, IsLargerThanOrEqualTo, 11)
c.Assert(22, Not(IsLargerThanOrEqualTo), 66)
c.Assert(22, IsLargerThanOrEqualTo, 22)
c.Assert(66.5, IsLargerThanOrEqualTo, 22.5)
c.Assert(11.5, Not(IsLargerThanOrEqualTo), 22.5)
c.Assert(22.5, IsLargerThanOrEqualTo, 22.5)
c.Assert(time.Duration(22), IsLargerThanOrEqualTo, time.Duration(11))
c.Assert(time.Duration(22), IsLargerThanOrEqualTo, time.Duration(22))
```

## Run tests of this library

This includes download of all dependencies and also creation and upload of coverage reports.

> No working golang installation is required but Java 8+ (in ``PATH`` or ``JAVA_HOME`` set.).

```bash
# On Linux/macOS
$ ./gradlew test

# On Windows
$ gradlew test
```

## Contributing

gocheck-addons is an open source project by [echocat](https://echocat.org).
So if you want to make this project even better, you can contribute to this project on [Github](https://github.com/echocat/gocheck-addons)
by [fork us](https://github.com/echocat/gocheck-addons/fork).

If you commit code to this project, you have to accept that this code will be released under the [license](#license) of this project.

## Support

If you need support you can create a ticket in our [issue tracker](https://github.com/echocat/gocheck-addons/issues).

## License

See the [LICENSE](LICENSE) file.
