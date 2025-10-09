//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_SessionManagementParameters) validateSetMethodParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SessionManagementParameters) validateSetParametersParameters(val ISessionManagementParametersParameters) error {
	return nil
}

func validateNewSessionManagementParametersParameters(options ISessionManagementParameters) error {
	return nil
}

