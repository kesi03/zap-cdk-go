//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IDataDrivenNode) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IDataDrivenNode) validateSetRegexParameters(val *string) error {
	return nil
}

