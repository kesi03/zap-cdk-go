//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateSOAPConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SOAPConfig) validateSetConfigParameters(val ISoap) error {
	return nil
}

func validateNewSOAPConfigParameters(scope constructs.Construct, id *string, props ISoapProps) error {
	return nil
}

