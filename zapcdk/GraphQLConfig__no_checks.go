//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateGraphQLConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GraphQLConfig) validateSetConfigParameters(val IGraphQL) error {
	return nil
}

func validateNewGraphQLConfigParameters(scope constructs.Construct, id *string, props IGraphQLProps) error {
	return nil
}

