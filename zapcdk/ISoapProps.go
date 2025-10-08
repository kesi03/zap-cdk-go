package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the SOAPConfig construct.
type ISoapProps interface {
	Soap() ISoap
	SetSoap(s ISoap)
}

// The jsii proxy for ISoapProps
type jsiiProxy_ISoapProps struct {
	_ byte // padding
}

func (j *jsiiProxy_ISoapProps) Soap() ISoap {
	var returns ISoap
	_jsii_.Get(
		j,
		"soap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoapProps)SetSoap(val ISoap) {
	if err := j.validateSetSoapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"soap",
		val,
	)
}

