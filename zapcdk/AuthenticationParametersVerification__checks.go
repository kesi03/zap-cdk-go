//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_AuthenticationParametersVerification) validateSetMethodParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAuthenticationParametersVerificationParameters(options IAuthenticationParametersVerification) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

