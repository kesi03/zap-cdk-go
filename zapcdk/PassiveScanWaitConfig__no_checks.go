//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validatePassiveScanWaitConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PassiveScanWaitConfig) validateSetConfigParameters(val IPassiveScanWait) error {
	return nil
}

func validateNewPassiveScanWaitConfigParameters(scope constructs.Construct, id *string, props IPassiveScanWaitProps) error {
	return nil
}

