//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ExitStatus) validateSetParametersParameters(val IExitStatusParameters) error {
	return nil
}

func (j *jsiiProxy_ExitStatus) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewExitStatusParameters(options IExitStatus) error {
	return nil
}

