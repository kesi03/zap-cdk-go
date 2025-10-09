//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Export) validateSetFileNameParameters(val *string) error {
	return nil
}

func validateNewExportParameters(options IExport) error {
	return nil
}

