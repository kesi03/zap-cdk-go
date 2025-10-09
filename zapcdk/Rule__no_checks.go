//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Rule) validateSetIdParameters(val *float64) error {
	return nil
}

func validateNewRuleParameters(options IRule) error {
	return nil
}

