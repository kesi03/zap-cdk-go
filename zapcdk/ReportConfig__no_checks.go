//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateReportConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ReportConfig) validateSetConfigParameters(val IReport) error {
	return nil
}

func validateNewReportConfigParameters(scope constructs.Construct, id *string, props IReportProps) error {
	return nil
}

