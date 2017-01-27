package addons

import (
	"gopkg.in/check.v1"
	"strings"
)

type contains struct {
	*check.CheckerInfo
}

var Contains check.Checker = &contains{
	&check.CheckerInfo{
		Name:   "Contains",
		Params: []string{"value", "contained"},
	},
}

func (checker *contains) checkString(value string, rawContained interface{}) (bool,  string) {
	if contained, ok := rawContained.(string); ok {
		return strings.Contains(value, contained), ""
	}
	if contained, ok := rawContained.(hasStringMethod); ok {
		return strings.Contains(value, contained.String()), ""
	}
	return false, "For string values only string contained values supported."
}

func (checker *contains) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		panic("Illegal number of parameters.")
	}
	if value, ok := params[0].(string); ok {
		return checker.checkString(value, params[1])
	}
	if value, ok := params[0].(hasStringMethod); ok {
		return checker.checkString(value.String(), params[1])
	}
	return false, "No compatible value."
}

type hasPrefix struct {
	*check.CheckerInfo
}

var HasPrefix check.Checker = &hasPrefix{
	&check.CheckerInfo{
		Name:   "HasPrefix",
		Params: []string{"value", "prefix"},
	},
}

func (checker *hasPrefix) checkString(value string, rawPrefix interface{}) (bool, string) {
	if prefix, ok := rawPrefix.(string); ok {
		return strings.HasPrefix(value, prefix), ""
	}
	if prefix, ok := rawPrefix.(hasStringMethod); ok {
		return strings.HasPrefix(value, prefix.String()), ""
	}
	return false, "For string values only string prefix values supported."
}

func (checker *hasPrefix) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		panic("Illegal number of parameters.")
	}
	if value, ok := params[0].(string); ok {
		return checker.checkString(value, params[1])
	}
	if value, ok := params[0].(hasStringMethod); ok {
		return checker.checkString(value.String(), params[1])
	}
	return false, "No compatible value."
}

type hasSuffix struct {
	*check.CheckerInfo
}

var HasSuffix check.Checker = &hasSuffix{
	&check.CheckerInfo{
		Name:   "HasSuffix",
		Params: []string{"value", "suffix"},
	},
}

func (checker *hasSuffix) checkString(value string, rawSuffix interface{}) (bool, string) {
	if suffix, ok := rawSuffix.(string); ok {
		return strings.HasSuffix(value, suffix), ""
	}
	if suffix, ok := rawSuffix.(hasStringMethod); ok {
		return strings.HasSuffix(value, suffix.String()), ""
	}
	return false, "For string values only string suffix values supported."
}

func (checker *hasSuffix) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		panic("Illegal number of parameters.")
	}
	if value, ok := params[0].(string); ok {
		return checker.checkString(value, params[1])
	}
	if value, ok := params[0].(hasStringMethod); ok {
		return checker.checkString(value.String(), params[1])
	}
	return false, "No compatible value."
}

