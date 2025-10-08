//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateExitStatusConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ExitStatusConfig) validateSetConfigParameters(val IExitStatus) error {
	return nil
}

func validateNewExitStatusConfigParameters(scope constructs.Construct, id *string, props IExitStatusProps) error {
	return nil
}

