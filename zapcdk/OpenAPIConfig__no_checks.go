//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateOpenAPIConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_OpenAPIConfig) validateSetConfigParameters(val IOpenAPI) error {
	return nil
}

func validateNewOpenAPIConfigParameters(scope constructs.Construct, id *string, props IOpenAPIProps) error {
	return nil
}

