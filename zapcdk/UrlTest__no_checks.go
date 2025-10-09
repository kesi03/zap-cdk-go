//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_UrlTest) validateSetOnFailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UrlTest) validateSetOperatorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UrlTest) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UrlTest) validateSetUrlParameters(val *string) error {
	return nil
}

func validateNewUrlTestParameters(options IUrlTest) error {
	return nil
}

