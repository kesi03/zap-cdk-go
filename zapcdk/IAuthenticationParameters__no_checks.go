//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IAuthenticationParameters) validateSetMethodParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IAuthenticationParameters) validateSetParametersParameters(val IAuthenticationParametersParameters) error {
	return nil
}

func (j *jsiiProxy_IAuthenticationParameters) validateSetVerificationParameters(val IAuthenticationParametersVerification) error {
	return nil
}

