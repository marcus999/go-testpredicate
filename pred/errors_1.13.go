// +build go1.13

package pred

import (
	"errors"
	"fmt"

	"github.com/maargenton/go-testpredicate"
)

// IsError tests if an error value matches a specific error,
// using errros.Is() to support go v1.13 error wrapping
func IsError(expectedErr error) testpredicate.Predicate {
	return testpredicate.MakePredicate(
		fmt.Sprintf("value is an error matching %v", expectedErr),
		func(value interface{}) (testpredicate.PredicateResult, error) {
			if value == nil {
				if expectedErr == nil {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed, nil
			} else if err, ok := value.(error); ok {
				if errors.Is(err, expectedErr) {
					return testpredicate.PredicatePassed, nil
				}
				return testpredicate.PredicateFailed,
					fmt.Errorf("detailed error: %v", err)
			} else {
				return testpredicate.PredicateInvalid,
					fmt.Errorf("value of type '%T' is not an error", value)
			}
		})
}