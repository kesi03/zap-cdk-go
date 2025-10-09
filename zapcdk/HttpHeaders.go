package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for HTTP header scanning.
//
// Example:
//   const headerConfig = new HttpHeaders({ enabled: false, allRequests: false });
//   console.log(headerConfig.enabled); // false
//
type HttpHeaders interface {
	IHttpHeaders
	// If headers of requests without parameters should be scanned.
	//
	// Default: false.
	AllRequests() *bool
	SetAllRequests(val *bool)
	// If HTTP header scanning should be enabled.
	//
	// Default: false.
	Enabled() *bool
	SetEnabled(val *bool)
}

// The jsii proxy struct for HttpHeaders
type jsiiProxy_HttpHeaders struct {
	jsiiProxy_IHttpHeaders
}

func (j *jsiiProxy_HttpHeaders) AllRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpHeaders) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}


// Creates an instance of HttpHeaders.
//
// Example:
//   const headerConfig = new HttpHeaders({ enabled: false, allRequests: false });
//   console.log(headerConfig.enabled); // false
//
func NewHttpHeaders(options IHttpHeaders) HttpHeaders {
	_init_.Initialize()

	j := jsiiProxy_HttpHeaders{}

	_jsii_.Create(
		"zap-cdk.HttpHeaders",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of HttpHeaders.
//
// Example:
//   const headerConfig = new HttpHeaders({ enabled: false, allRequests: false });
//   console.log(headerConfig.enabled); // false
//
func NewHttpHeaders_Override(h HttpHeaders, options IHttpHeaders) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.HttpHeaders",
		[]interface{}{options},
		h,
	)
}

func (j *jsiiProxy_HttpHeaders)SetAllRequests(val *bool) {
	_jsii_.Set(
		j,
		"allRequests",
		val,
	)
}

func (j *jsiiProxy_HttpHeaders)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

