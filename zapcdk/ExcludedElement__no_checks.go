//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_ExcludedElement) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ExcludedElement) validateSetElementParameters(val *string) error {
	return nil
}

func validateNewExcludedElementParameters(options IExcludedElement) error {
	return nil
}

