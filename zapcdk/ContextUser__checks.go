//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_ContextUser) validateSetCredentialsParameters(val *[]IUserCredentials) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ContextUser) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewContextUserParameters(options IContextUser) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

