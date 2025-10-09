package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration options for scanning URL query strings and Data Driven Nodes (DDNs).
//
// Example:
//   const config = new UrlQueryStringAndDataDrivenNodes({ enabled: true, addParam: false, odata: true });
//   console.log(config.enabled); // true
//
type UrlQueryStringAndDataDrivenNodes interface {
	IUrlQueryStringAndDataDrivenNodes
	// If a query parameter should be added if none present.
	//
	// Default: false.
	AddParam() *bool
	SetAddParam(val *bool)
	// If query parameters and DDNs scanning should be enabled.
	//
	// Default: true.
	Enabled() *bool
	SetEnabled(val *bool)
	// If OData query filters should be scanned.
	//
	// Default: true.
	Odata() *bool
	SetOdata(val *bool)
}

// The jsii proxy struct for UrlQueryStringAndDataDrivenNodes
type jsiiProxy_UrlQueryStringAndDataDrivenNodes struct {
	jsiiProxy_IUrlQueryStringAndDataDrivenNodes
}

func (j *jsiiProxy_UrlQueryStringAndDataDrivenNodes) AddParam() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlQueryStringAndDataDrivenNodes) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlQueryStringAndDataDrivenNodes) Odata() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"odata",
		&returns,
	)
	return returns
}


// Creates an instance of UrlQueryStringAndDataDrivenNodes.
//
// Example:
//   const config = new UrlQueryStringAndDataDrivenNodes({ enabled: true, addParam: false, odata: true });
//   console.log(config.enabled); // true
//
func NewUrlQueryStringAndDataDrivenNodes(options IUrlQueryStringAndDataDrivenNodes) UrlQueryStringAndDataDrivenNodes {
	_init_.Initialize()

	j := jsiiProxy_UrlQueryStringAndDataDrivenNodes{}

	_jsii_.Create(
		"zap-cdk.UrlQueryStringAndDataDrivenNodes",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of UrlQueryStringAndDataDrivenNodes.
//
// Example:
//   const config = new UrlQueryStringAndDataDrivenNodes({ enabled: true, addParam: false, odata: true });
//   console.log(config.enabled); // true
//
func NewUrlQueryStringAndDataDrivenNodes_Override(u UrlQueryStringAndDataDrivenNodes, options IUrlQueryStringAndDataDrivenNodes) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.UrlQueryStringAndDataDrivenNodes",
		[]interface{}{options},
		u,
	)
}

func (j *jsiiProxy_UrlQueryStringAndDataDrivenNodes)SetAddParam(val *bool) {
	_jsii_.Set(
		j,
		"addParam",
		val,
	)
}

func (j *jsiiProxy_UrlQueryStringAndDataDrivenNodes)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_UrlQueryStringAndDataDrivenNodes)SetOdata(val *bool) {
	_jsii_.Set(
		j,
		"odata",
		val,
	)
}

