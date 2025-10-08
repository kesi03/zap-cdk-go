//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validatePostmanConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_PostmanConfig) validateSetConfigParameters(val IPostman) error {
	return nil
}

func validateNewPostmanConfigParameters(scope constructs.Construct, id *string, props IPostmanProps) error {
	return nil
}

