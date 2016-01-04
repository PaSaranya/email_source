package matchers

import (
	"fmt"
	"reflect"

	"github.com/pivotal-cf/email-resource/Godeps/_workspace/src/github.com/onsi/gomega/format"
)

type AssignableToTypeOfMatcher struct {
	Expected interface{}
}

func (matcher *AssignableToTypeOfMatcher) Match(actual interface{}) (success bool, err error) {
	if actual == nil || matcher.Expected == nil {
		return false, fmt.Errorf("Refusing to compare <nil> to <nil>.\nBe explicit and use BeNil() instead.  This is to avoid mistakes where both sides of an assertion are erroneously uninitialized.")
	}

	actualType := reflect.TypeOf(actual)
	expectedType := reflect.TypeOf(matcher.Expected)

	return actualType.AssignableTo(expectedType), nil
}

func (matcher *AssignableToTypeOfMatcher) FailureMessage(actual interface{}) string {
	return format.Message(actual, fmt.Sprintf("to be assignable to the type: %T", matcher.Expected))
}

func (matcher *AssignableToTypeOfMatcher) NegatedFailureMessage(actual interface{}) string {
	return format.Message(actual, fmt.Sprintf("not to be assignable to the type: %T", matcher.Expected))
}
