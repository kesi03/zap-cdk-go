//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IReplacerRule) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IReplacerRule) validateSetMatchStringParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IReplacerRule) validateSetMatchTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IReplacerRule) validateSetReplacementStringParameters(val *string) error {
	return nil
}

