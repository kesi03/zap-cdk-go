//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateExportConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ExportConfig) validateSetConfigParameters(val IExport) error {
	return nil
}

func validateNewExportConfigParameters(scope constructs.Construct, id *string, props IExportProps) error {
	return nil
}

