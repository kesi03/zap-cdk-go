//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IContext) validateSetAuthenticationParameters(val IAuthenticationParameters) error {
	return nil
}

func (j *jsiiProxy_IContext) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IContext) validateSetSessionManagementParameters(val ISessionManagementParameters) error {
	return nil
}

func (j *jsiiProxy_IContext) validateSetStructureParameters(val IContextStructure) error {
	return nil
}

func (j *jsiiProxy_IContext) validateSetTechnologyParameters(val ITechnology) error {
	return nil
}

func (j *jsiiProxy_IContext) validateSetUrlsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_IContext) validateSetUsersParameters(val *[]IContextUser) error {
	return nil
}

