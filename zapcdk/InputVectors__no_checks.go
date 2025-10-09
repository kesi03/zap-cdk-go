//go:build no_runtime_type_checking

package zapcdk

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_InputVectors) validateSetCookieDataParameters(val ICookieData) error {
	return nil
}

func (j *jsiiProxy_InputVectors) validateSetHttpHeadersParameters(val IHttpHeaders) error {
	return nil
}

func (j *jsiiProxy_InputVectors) validateSetPostDataParameters(val IPostData) error {
	return nil
}

func (j *jsiiProxy_InputVectors) validateSetUrlQueryStringAndDataDrivenNodesParameters(val IUrlQueryStringAndDataDrivenNodes) error {
	return nil
}

