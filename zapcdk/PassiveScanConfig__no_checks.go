//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_PassiveScanConfig) validateSetParametersParameters(val IPassiveScanParameters) error {
	return nil
}

func (j *jsiiProxy_PassiveScanConfig) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewPassiveScanConfigParameters(options IPassiveScanConfig) error {
	return nil
}

