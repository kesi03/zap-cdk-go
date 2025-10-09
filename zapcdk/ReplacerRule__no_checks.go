//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ReplacerRule) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplacerRule) validateSetMatchStringParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplacerRule) validateSetMatchTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ReplacerRule) validateSetReplacementStringParameters(val *string) error {
	return nil
}

func validateNewReplacerRuleParameters(options IReplacerRule) error {
	return nil
}

