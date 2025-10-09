//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validatePassiveScanConfigConstruct_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PassiveScanConfigConstruct) validateSetConfigParameters(val IPassiveScanConfig) error {
	return nil
}

func validateNewPassiveScanConfigConstructParameters(scope constructs.Construct, id *string, props IPassiveScanConfigProps) error {
	return nil
}

