//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validatePassiveScanConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PassiveScanConfig) validateSetConfigParameters(val IPassiveScanConfig) error {
	return nil
}

func validateNewPassiveScanConfigParameters(scope constructs.Construct, id *string, props IPassiveScanConfigProps) error {
	return nil
}

