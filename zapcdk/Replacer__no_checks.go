//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Replacer) validateSetRulesParameters(val *[]IReplacerRule) error {
	return nil
}

func validateNewReplacerParameters(options IReplacer) error {
	return nil
}

