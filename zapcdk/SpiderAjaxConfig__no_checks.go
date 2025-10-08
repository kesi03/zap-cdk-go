//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func validateSpiderAjaxConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_SpiderAjaxConfig) validateSetConfigParameters(val ISpiderAjax) error {
	return nil
}

func validateNewSpiderAjaxConfigParameters(scope constructs.Construct, id *string, props ISpiderAjaxProps) error {
	return nil
}

