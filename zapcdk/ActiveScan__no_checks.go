//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ActiveScan) validateSetParametersParameters(val IActiveScanParameters) error {
	return nil
}

func (j *jsiiProxy_ActiveScan) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewActiveScanParameters(options IActiveScan) error {
	return nil
}

