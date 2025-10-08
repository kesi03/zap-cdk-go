//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validatePostmanConfig_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_PostmanConfig) validateSetConfigParameters(val IPostman) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewPostmanConfigParameters(scope constructs.Construct, id *string, props IPostmanProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

