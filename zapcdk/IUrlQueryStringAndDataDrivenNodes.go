package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration options for scanning URL query strings and Data Driven Nodes (DDNs).
type IUrlQueryStringAndDataDrivenNodes interface {
	// If a query parameter should be added if none present.
	//
	// Default: false.
	AddParam() *bool
	SetAddParam(a *bool)
	// If query parameters and DDNs scanning should be enabled.
	//
	// Default: true.
	Enabled() *bool
	SetEnabled(e *bool)
	// If OData query filters should be scanned.
	//
	// Default: true.
	Odata() *bool
	SetOdata(o *bool)
}

// The jsii proxy for IUrlQueryStringAndDataDrivenNodes
type jsiiProxy_IUrlQueryStringAndDataDrivenNodes struct {
	_ byte // padding
}

func (j *jsiiProxy_IUrlQueryStringAndDataDrivenNodes) AddParam() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlQueryStringAndDataDrivenNodes)SetAddParam(val *bool) {
	_jsii_.Set(
		j,
		"addParam",
		val,
	)
}

func (j *jsiiProxy_IUrlQueryStringAndDataDrivenNodes) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlQueryStringAndDataDrivenNodes)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IUrlQueryStringAndDataDrivenNodes) Odata() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"odata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlQueryStringAndDataDrivenNodes)SetOdata(val *bool) {
	_jsii_.Set(
		j,
		"odata",
		val,
	)
}

