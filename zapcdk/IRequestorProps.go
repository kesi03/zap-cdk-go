package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the RequestorConfig construct.
type IRequestorProps interface {
	Requestor() IRequestorParameters
	SetRequestor(r IRequestorParameters)
}

// The jsii proxy for IRequestorProps
type jsiiProxy_IRequestorProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IRequestorProps) Requestor() IRequestorParameters {
	var returns IRequestorParameters
	_jsii_.Get(
		j,
		"requestor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestorProps)SetRequestor(val IRequestorParameters) {
	if err := j.validateSetRequestorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestor",
		val,
	)
}

