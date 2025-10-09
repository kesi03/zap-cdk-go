//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_InputVectors) validateSetCookieDataParameters(val ICookieData) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_InputVectors) validateSetHttpHeadersParameters(val IHttpHeaders) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_InputVectors) validateSetPostDataParameters(val IPostData) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_InputVectors) validateSetUrlQueryStringAndDataDrivenNodesParameters(val IUrlQueryStringAndDataDrivenNodes) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

