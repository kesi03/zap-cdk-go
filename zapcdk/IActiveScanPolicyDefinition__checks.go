//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
	"time"
)

func (j *jsiiProxy_IActiveScanPolicyDefinition) validateSetCreatedAtParameters(val *time.Time) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) validateSetIdParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IActiveScanPolicyDefinition) validateSetUpdatedAtParameters(val *time.Time) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

