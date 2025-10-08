//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateDelayConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_DelayConfig) validateSetConfigParameters(val IDelay) error {
	return nil
}

func validateNewDelayConfigParameters(scope constructs.Construct, id *string, props IDelayProps) error {
	return nil
}

