//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_AuthenticationParameters) validateSetMethodParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AuthenticationParameters) validateSetParametersParameters(val IAuthenticationParametersParameters) error {
	return nil
}

func (j *jsiiProxy_AuthenticationParameters) validateSetVerificationParameters(val IAuthenticationParametersVerification) error {
	return nil
}

func validateNewAuthenticationParametersParameters(options IAuthenticationParameters) error {
	return nil
}

