//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_IInputVectors) validateSetCookieDataParameters(val ICookieData) error {
	return nil
}

func (j *jsiiProxy_IInputVectors) validateSetHttpHeadersParameters(val IHttpHeaders) error {
	return nil
}

func (j *jsiiProxy_IInputVectors) validateSetPostDataParameters(val IPostData) error {
	return nil
}

func (j *jsiiProxy_IInputVectors) validateSetUrlQueryStringAndDataDrivenNodesParameters(val IUrlQueryStringAndDataDrivenNodes) error {
	return nil
}

