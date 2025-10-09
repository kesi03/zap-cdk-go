package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the SOAP service configuration.
//
// Example:
//   const soapConfig = new Soap({
//     wsdlFile: 'service.wsdl',
//     wsdlUrl: 'https://example.com/service?wsdl'
//   });
//
type Soap interface {
	ISoap
	WsdlFile() *string
	SetWsdlFile(val *string)
	WsdlUrl() *string
	SetWsdlUrl(val *string)
}

// The jsii proxy struct for Soap
type jsiiProxy_Soap struct {
	jsiiProxy_ISoap
}

func (j *jsiiProxy_Soap) WsdlFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wsdlFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Soap) WsdlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wsdlUrl",
		&returns,
	)
	return returns
}


// Creates an instance of Soap.
func NewSoap(options ISoap) Soap {
	_init_.Initialize()

	if err := validateNewSoapParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Soap{}

	_jsii_.Create(
		"zap-cdk.Soap",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Soap.
func NewSoap_Override(s Soap, options ISoap) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Soap",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_Soap)SetWsdlFile(val *string) {
	_jsii_.Set(
		j,
		"wsdlFile",
		val,
	)
}

func (j *jsiiProxy_Soap)SetWsdlUrl(val *string) {
	_jsii_.Set(
		j,
		"wsdlUrl",
		val,
	)
}

