//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IReplacer) validateSetRulesParameters(val *[]IReplacerRule) error {
	return nil
}

