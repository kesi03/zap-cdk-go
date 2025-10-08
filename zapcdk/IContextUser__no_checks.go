//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IContextUser) validateSetCredentialsParameters(val *[]IUserCredentials) error {
	return nil
}

func (j *jsiiProxy_IContextUser) validateSetNameParameters(val *string) error {
	return nil
}

