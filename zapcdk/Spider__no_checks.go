//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Spider) validateSetParametersParameters(val ISpiderParameters) error {
	return nil
}

func (j *jsiiProxy_Spider) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewSpiderParameters(parameters ISpiderParameters) error {
	return nil
}

