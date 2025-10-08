//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_IAuthenticationParameters) validateSetMethodParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IAuthenticationParameters) validateSetParametersParameters(val IAuthenticationParametersParameters) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IAuthenticationParameters) validateSetVerificationParameters(val IAuthenticationParametersVerification) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

