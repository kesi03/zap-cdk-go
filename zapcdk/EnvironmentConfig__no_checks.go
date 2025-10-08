//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_EnvironmentConfig) validateSetConfigParameters(val IEnvironment) error {
	return nil
}

func validateNewEnvironmentConfigParameters(scope constructs.Construct, id *string, props IEnvironmentProps) error {
	return nil
}

