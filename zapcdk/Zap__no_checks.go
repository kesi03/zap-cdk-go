//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Zap) validateSetEnvParameters(val IEnvironment) error {
	return nil
}

func (j *jsiiProxy_Zap) validateSetJobsParameters(val *[]interface{}) error {
	return nil
}

func validateNewZapParameters(options IZap) error {
	return nil
}

