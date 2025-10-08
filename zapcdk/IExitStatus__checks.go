//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_IExitStatus) validateSetParametersParameters(val IExitStatusParameters) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IExitStatus) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

