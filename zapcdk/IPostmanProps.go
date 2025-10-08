package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the PostmanConfig construct.
type IPostmanProps interface {
	Postman() IPostman
	SetPostman(p IPostman)
}

// The jsii proxy for IPostmanProps
type jsiiProxy_IPostmanProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IPostmanProps) Postman() IPostman {
	var returns IPostman
	_jsii_.Get(
		j,
		"postman",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostmanProps)SetPostman(val IPostman) {
	if err := j.validateSetPostmanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postman",
		val,
	)
}

