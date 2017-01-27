package addons

import (
	"fmt"
	"gopkg.in/check.v1"
	"reflect"
	"regexp"
)

type hasStringMethod interface {
	String() string
}

type throwsPanicThatMatches struct {
	*check.CheckerInfo
}

var ThrowsPanicThatMatches check.Checker = &throwsPanicThatMatches{
	&check.CheckerInfo{
		Name:   "ThrowsPanicThatMatches",
		Params: []string{"action", "regex"},
	},
}

func (checker *throwsPanicThatMatches) Check(params []interface{}, names []string) (result bool, errorMessage string) {
	if len(params) != 2 {
		panic("Illegal number of parameters.")
	}
	action, ok := params[0].(func())
	if !ok {
		return false, "There is no <action> function with signature func() provided."
	}
	plainRegex, ok := params[1].(string)
	if !ok {
		return false, "There is no <regex> of type string provided."
	}
	regex, err := regexp.Compile(plainRegex)
	if err != nil {
		return false, fmt.Sprintf("Can't compile regex '%s'. Got: %s", plainRegex, err.Error())
	}
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				params[0] = err.Error()
				result = regex.MatchString(err.Error())
			} else if rs, ok := r.(hasStringMethod); ok {
				params[0] = rs.String()
				result = regex.MatchString(rs.String())
			} else if rs, ok := r.(string); ok {
				params[0] = rs
				result = regex.MatchString(rs)
			} else {
				params[0] = r
				result = false
			}
		}
	}()
	action()
	return false, ""
}

type isEmptyChecker struct {
	*check.CheckerInfo
}

var IsEmpty check.Checker = &isEmptyChecker{
	&check.CheckerInfo{
		Name:   "IsEmpty",
		Params: []string{"value"},
	},
}

func (checker *isEmptyChecker) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 1 {
		panic("Illegal number of parameters.")
	}
	param := params[0]
	if asString, ok := param.(string); ok {
		return len(asString) == 0, ""
	}
	pv := reflect.ValueOf(param)
	switch pv.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		return pv.Len() == 0, ""
	}
	return false, "No compatible value."
}
