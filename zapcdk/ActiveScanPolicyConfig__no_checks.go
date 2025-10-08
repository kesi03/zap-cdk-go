//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateActiveScanPolicyConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ActiveScanPolicyConfig) validateSetConfigParameters(val IActiveScanPolicy) error {
	return nil
}

func validateNewActiveScanPolicyConfigParameters(scope constructs.Construct, id *string, props IActiveScanPolicyProps) error {
	return nil
}

