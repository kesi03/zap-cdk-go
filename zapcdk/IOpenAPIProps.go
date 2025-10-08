package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the OpenAPIConfig construct.
type IOpenAPIProps interface {
	Openapi() IOpenAPI
	SetOpenapi(o IOpenAPI)
}

// The jsii proxy for IOpenAPIProps
type jsiiProxy_IOpenAPIProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IOpenAPIProps) Openapi() IOpenAPI {
	var returns IOpenAPI
	_jsii_.Get(
		j,
		"openapi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenAPIProps)SetOpenapi(val IOpenAPI) {
	if err := j.validateSetOpenapiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openapi",
		val,
	)
}

