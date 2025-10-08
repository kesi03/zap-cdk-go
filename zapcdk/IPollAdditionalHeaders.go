package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPollAdditionalHeaders interface {
	Header() *string
	SetHeader(h *string)
	Value() *string
	SetValue(v *string)
}

// The jsii proxy for IPollAdditionalHeaders
type jsiiProxy_IPollAdditionalHeaders struct {
	_ byte // padding
}

func (j *jsiiProxy_IPollAdditionalHeaders) Header() *string {
	var returns *string
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPollAdditionalHeaders)SetHeader(val *string) {
	if err := j.validateSetHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"header",
		val,
	)
}

func (j *jsiiProxy_IPollAdditionalHeaders) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPollAdditionalHeaders)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

