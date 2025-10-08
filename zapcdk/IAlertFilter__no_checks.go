//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IAlertFilter) validateSetNewRiskParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IAlertFilter) validateSetRuleIdParameters(val *float64) error {
	return nil
}

