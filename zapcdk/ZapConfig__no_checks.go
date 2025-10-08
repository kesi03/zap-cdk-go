//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateZapConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ZapConfig) validateSetConfigParameters(val IZap) error {
	return nil
}

func validateNewZapConfigParameters(scope constructs.Construct, id *string, props IZap) error {
	return nil
}

