package impl_test

import (
	"fmt"
	"testing"

	"github.com/maargenton/go-testpredicate/pkg/internal/predicate"
	"github.com/maargenton/go-testpredicate/pkg/internal/predicate/impl"
)

func TestIs(t *testing.T) {
	verifyPredicate(t, pr(impl.Is(customPredicate())), expectation{value: nil, pass: true})
	verifyPredicate(t, pr(impl.Is(customPredicate())), expectation{value: 123, pass: false})
}

func customPredicate() (desc string, f predicate.PredicateFunc) {
	desc = "{} is custom"
	f = func(value interface{}) (success bool, ctx []predicate.ContextValue, err error) {
		success = value == nil
		return
	}
	return
}

func TestEval(t *testing.T) {
	verifyTransform(t, tr(impl.Eval(customTransform())), expectation{
		value:  123,
		result: "123",
	})
}

func customTransform() (desc string, f predicate.TransformFunc) {
	desc = "custom({})"
	f = func(value interface{}) (result interface{}, ctx []predicate.ContextValue, err error) {
		result = fmt.Sprintf("%v", value)
		return
	}
	return
}
