package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for JSON body scanning in POST data.
type IJsonPostData interface {
	// If JSON scanning should be enabled.
	//
	// Default: true.
	Enabled() *bool
	SetEnabled(e *bool)
	// If null values should be scanned.
	//
	// Default: false.
	ScanNullValues() *bool
	SetScanNullValues(s *bool)
}

// The jsii proxy for IJsonPostData
type jsiiProxy_IJsonPostData struct {
	_ byte // padding
}

func (j *jsiiProxy_IJsonPostData) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonPostData)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IJsonPostData) ScanNullValues() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanNullValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJsonPostData)SetScanNullValues(val *bool) {
	_jsii_.Set(
		j,
		"scanNullValues",
		val,
	)
}

