//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateActiveScanConfigConstruct_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ActiveScanConfigConstruct) validateSetConfigParameters(val IActiveScanConfig) error {
	return nil
}

func validateNewActiveScanConfigConstructParameters(scope constructs.Construct, id *string, props IActiveScanConfigProps) error {
	return nil
}

