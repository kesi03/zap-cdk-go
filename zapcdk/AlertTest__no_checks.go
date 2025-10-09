//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_AlertTest) validateSetOnFailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertTest) validateSetScanRuleIdParameters(val *float64) error {
	return nil
}

func (j *jsiiProxy_AlertTest) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewAlertTestParameters(options IAlertTest) error {
	return nil
}

