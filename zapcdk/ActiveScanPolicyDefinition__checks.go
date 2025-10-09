//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
	"time"
)

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetCreatedAtParameters(val *time.Time) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetIdParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetUpdatedAtParameters(val *time.Time) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewActiveScanPolicyDefinitionParameters(options IActiveScanPolicyDefinition) error {
	if options == nil {
		return fmt.Errorf("parameter options is required, but nil was provided")
	}

	return nil
}

