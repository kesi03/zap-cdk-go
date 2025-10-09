//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_SpiderTest) validateSetOnFailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpiderTest) validateSetOperatorParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpiderTest) validateSetStatisticParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpiderTest) validateSetTypeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpiderTest) validateSetValueParameters(val *float64) error {
	return nil
}

func validateNewSpiderTestParameters(options ISpiderTest) error {
	return nil
}

