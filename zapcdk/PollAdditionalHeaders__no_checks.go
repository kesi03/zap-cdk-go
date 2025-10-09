//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_PollAdditionalHeaders) validateSetHeaderParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PollAdditionalHeaders) validateSetValueParameters(val *string) error {
	return nil
}

func validateNewPollAdditionalHeadersParameters(options IPollAdditionalHeaders) error {
	return nil
}

