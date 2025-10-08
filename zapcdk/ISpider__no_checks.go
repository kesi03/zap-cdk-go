//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ISpider) validateSetParametersParameters(val ISpiderParameters) error {
	return nil
}

func (j *jsiiProxy_ISpider) validateSetTypeParameters(val *string) error {
	return nil
}

