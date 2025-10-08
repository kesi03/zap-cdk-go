package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for POST data scanning.
type IPostData interface {
	// If DWR scanning should be enabled.
	//
	// Default: false.
	DirectWebRemoting() *bool
	SetDirectWebRemoting(d *bool)
	// If POST data scanning is enabled.
	//
	// Default: true.
	Enabled() *bool
	SetEnabled(e *bool)
	// If GWT scanning should be enabled.
	//
	// Default: false.
	GoogleWebToolkit() *bool
	SetGoogleWebToolkit(g *bool)
	// Configuration for JSON bodies.
	Json() IJsonPostData
	SetJson(j IJsonPostData)
	// If multipart form data bodies should be scanned.
	//
	// Default: true.
	MultiPartFormData() *bool
	SetMultiPartFormData(m *bool)
	// If XML bodies should be scanned.
	//
	// Default: true.
	Xml() *bool
	SetXml(x *bool)
}

// The jsii proxy for IPostData
type jsiiProxy_IPostData struct {
	_ byte // padding
}

func (j *jsiiProxy_IPostData) DirectWebRemoting() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"directWebRemoting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostData)SetDirectWebRemoting(val *bool) {
	_jsii_.Set(
		j,
		"directWebRemoting",
		val,
	)
}

func (j *jsiiProxy_IPostData) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostData)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IPostData) GoogleWebToolkit() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"googleWebToolkit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostData)SetGoogleWebToolkit(val *bool) {
	_jsii_.Set(
		j,
		"googleWebToolkit",
		val,
	)
}

func (j *jsiiProxy_IPostData) Json() IJsonPostData {
	var returns IJsonPostData
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostData)SetJson(val IJsonPostData) {
	if err := j.validateSetJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_IPostData) MultiPartFormData() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multiPartFormData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostData)SetMultiPartFormData(val *bool) {
	_jsii_.Set(
		j,
		"multiPartFormData",
		val,
	)
}

func (j *jsiiProxy_IPostData) Xml() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostData)SetXml(val *bool) {
	_jsii_.Set(
		j,
		"xml",
		val,
	)
}

