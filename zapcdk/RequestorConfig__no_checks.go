//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateRequestorConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RequestorConfig) validateSetConfigParameters(val IRequestorParameters) error {
	return nil
}

func validateNewRequestorConfigParameters(scope constructs.Construct, id *string, props IRequestorProps) error {
	return nil
}

