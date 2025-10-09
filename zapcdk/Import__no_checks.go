//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Import) validateSetFileNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Import) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewImportParameters(options IImport) error {
	return nil
}

