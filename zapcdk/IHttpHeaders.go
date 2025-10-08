package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for HTTP header scanning.
type IHttpHeaders interface {
	// If headers of requests without parameters should be scanned.
	//
	// Default: false.
	AllRequests() *bool
	SetAllRequests(a *bool)
	// If HTTP header scanning should be enabled.
	//
	// Default: false.
	Enabled() *bool
	SetEnabled(e *bool)
}

// The jsii proxy for IHttpHeaders
type jsiiProxy_IHttpHeaders struct {
	_ byte // padding
}

func (j *jsiiProxy_IHttpHeaders) AllRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpHeaders)SetAllRequests(val *bool) {
	_jsii_.Set(
		j,
		"allRequests",
		val,
	)
}

func (j *jsiiProxy_IHttpHeaders) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHttpHeaders)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

