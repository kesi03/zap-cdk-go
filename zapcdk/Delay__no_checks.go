//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Delay) validateSetParametersParameters(val IDelayParameters) error {
	return nil
}

func (j *jsiiProxy_Delay) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewDelayParameters(parameters IDelayParameters) error {
	return nil
}

