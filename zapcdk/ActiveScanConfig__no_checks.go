//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ActiveScanConfig) validateSetParametersParameters(val IActiveScanConfigParameters) error {
	return nil
}

func (j *jsiiProxy_ActiveScanConfig) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewActiveScanConfigParameters(options IActiveScanConfig) error {
	return nil
}

