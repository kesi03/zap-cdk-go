//go:build !no_runtime_type_checking

package zapcdk

import (
	"fmt"
)

func (j *jsiiProxy_IInputVectors) validateSetCookieDataParameters(val ICookieData) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IInputVectors) validateSetHttpHeadersParameters(val IHttpHeaders) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IInputVectors) validateSetPostDataParameters(val IPostData) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_IInputVectors) validateSetUrlQueryStringAndDataDrivenNodesParameters(val IUrlQueryStringAndDataDrivenNodes) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

