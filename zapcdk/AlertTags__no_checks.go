//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_AlertTags) validateSetExcludeParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_AlertTags) validateSetIncludeParameters(val *[]*string) error {
	return nil
}

func validateNewAlertTagsParameters(options IAlertTags) error {
	return nil
}

