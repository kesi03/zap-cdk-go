//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ActiveScanPolicy) validateSetParametersParameters(val IActiveScanPolicyParameters) error {
	return nil
}

func (j *jsiiProxy_ActiveScanPolicy) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewActiveScanPolicyParameters(options IActiveScanPolicy) error {
	return nil
}

