//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IEnvironment) validateSetContextsParameters(val *[]IContext) error {
	return nil
}

func (j *jsiiProxy_IEnvironment) validateSetParametersParameters(val IEnvironmentParameters) error {
	return nil
}

