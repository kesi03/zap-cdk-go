//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetCreatedAtParameters(val *time.Time) error {
	return nil
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActiveScanPolicyDefinition) validateSetUpdatedAtParameters(val *time.Time) error {
	return nil
}

func validateNewActiveScanPolicyDefinitionParameters(options IActiveScanPolicyDefinition) error {
	return nil
}

