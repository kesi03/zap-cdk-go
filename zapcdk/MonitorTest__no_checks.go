//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_MonitorTest) validateSetOnFailParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorTest) validateSetStatisticParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MonitorTest) validateSetTypeParameters(val *string) error {
	return nil
}

func validateNewMonitorTestParameters(options IMonitorTest) error {
	return nil
}

