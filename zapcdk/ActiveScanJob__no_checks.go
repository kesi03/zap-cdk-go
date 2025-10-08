//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateActiveScanJob_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ActiveScanJob) validateSetJobParameters(val IActiveScanJob) error {
	return nil
}

func validateNewActiveScanJobParameters(scope constructs.Construct, id *string, props IActiveScanJob) error {
	return nil
}

