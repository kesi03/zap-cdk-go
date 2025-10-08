//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateReplacerConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ReplacerConfig) validateSetConfigParameters(val IReplacer) error {
	return nil
}

func validateNewReplacerConfigParameters(scope constructs.Construct, id *string, props IReplacerProps) error {
	return nil
}

