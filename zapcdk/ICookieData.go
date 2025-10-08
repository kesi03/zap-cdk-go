package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for cookie data scanning.
type ICookieData interface {
	// If cookie scanning is enabled.
	//
	// Default: false.
	Enabled() *bool
	SetEnabled(e *bool)
	// If cookie values should be encoded.
	//
	// Default: false.
	EncodeCookieValues() *bool
	SetEncodeCookieValues(e *bool)
}

// The jsii proxy for ICookieData
type jsiiProxy_ICookieData struct {
	_ byte // padding
}

func (j *jsiiProxy_ICookieData) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICookieData)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ICookieData) EncodeCookieValues() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"encodeCookieValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICookieData)SetEncodeCookieValues(val *bool) {
	_jsii_.Set(
		j,
		"encodeCookieValues",
		val,
	)
}

