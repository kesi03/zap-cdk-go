//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_ISessionManagementParameters) validateSetMethodParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ISessionManagementParameters) validateSetParametersParameters(val ISessionManagementParametersParameters) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

