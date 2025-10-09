//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Rules) validateSetIdParameters(val *float64) error {
	return nil
}

func validateNewRulesParameters(options IRules) error {
	return nil
}

