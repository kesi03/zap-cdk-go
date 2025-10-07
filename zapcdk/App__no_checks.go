//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateApp_IsConstructParameters(x interface{}) error {
	return nil
}

