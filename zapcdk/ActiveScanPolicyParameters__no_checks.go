//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ActiveScanPolicyParameters) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActiveScanPolicyParameters) validateSetPolicyDefinitionParameters(val IActiveScanPolicyDefinition) error {
	return nil
}

func validateNewActiveScanPolicyParametersParameters(options IActiveScanPolicyParameters) error {
	return nil
}

