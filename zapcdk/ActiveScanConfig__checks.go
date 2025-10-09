//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_ActiveScanConfig) validateSetParametersParameters(val IActiveScanConfigParameters) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ActiveScanConfig) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewActiveScanConfigParameters(options IActiveScanConfig) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

