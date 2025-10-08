//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateSpiderConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SpiderConfig) validateSetConfigParameters(val ISpider) error {
	return nil
}

func validateNewSpiderConfigParameters(scope constructs.Construct, id *string, props *SpiderProps) error {
	return nil
}

