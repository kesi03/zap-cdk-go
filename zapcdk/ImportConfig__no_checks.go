//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateImportConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ImportConfig) validateSetConfigParameters(val IImport) error {
	return nil
}

func validateNewImportConfigParameters(scope constructs.Construct, id *string, props IImportProps) error {
	return nil
}

