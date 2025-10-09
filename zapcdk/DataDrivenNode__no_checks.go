//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_DataDrivenNode) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataDrivenNode) validateSetRegexParameters(val *string) error {
	return nil
}

func validateNewDataDrivenNodeParameters(options IDataDrivenNode) error {
	return nil
}

