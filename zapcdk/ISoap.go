package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing the configuration for a SOAP service.
type ISoap interface {
	WsdlFile() *string
	SetWsdlFile(w *string)
	WsdlUrl() *string
	SetWsdlUrl(w *string)
}

// The jsii proxy for ISoap
type jsiiProxy_ISoap struct {
	_ byte // padding
}

func (j *jsiiProxy_ISoap) WsdlFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wsdlFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoap)SetWsdlFile(val *string) {
	_jsii_.Set(
		j,
		"wsdlFile",
		val,
	)
}

func (j *jsiiProxy_ISoap) WsdlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wsdlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISoap)SetWsdlUrl(val *string) {
	_jsii_.Set(
		j,
		"wsdlUrl",
		val,
	)
}

