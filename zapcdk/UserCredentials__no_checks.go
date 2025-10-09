//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_UserCredentials) validateSetPasswordParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserCredentials) validateSetUsernameParameters(val *string) error {
	return nil
}

func validateNewUserCredentialsParameters(options IUserCredentials) error {
	return nil
}

