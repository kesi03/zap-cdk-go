package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for POST data scanning.
//
// Example:
//   const postDataConfig = new PostData({ enabled: true, multiPartFormData: true, xml: true, json: new JsonPostData(), googleWebToolkit: false, directWebRemoting: false });
//   console.log(postDataConfig.enabled); // true
//
type PostData interface {
	IPostData
	// If DWR scanning should be enabled.
	//
	// Default: false.
	DirectWebRemoting() *bool
	SetDirectWebRemoting(val *bool)
	// If POST data scanning is enabled.
	//
	// Default: true.
	Enabled() *bool
	SetEnabled(val *bool)
	// If GWT scanning should be enabled.
	//
	// Default: false.
	GoogleWebToolkit() *bool
	SetGoogleWebToolkit(val *bool)
	// Configuration for JSON bodies.
	Json() IJsonPostData
	SetJson(val IJsonPostData)
	// If multipart form data bodies should be scanned.
	//
	// Default: true.
	MultiPartFormData() *bool
	SetMultiPartFormData(val *bool)
	// If XML bodies should be scanned.
	//
	// Default: true.
	Xml() *bool
	SetXml(val *bool)
}

// The jsii proxy struct for PostData
type jsiiProxy_PostData struct {
	jsiiProxy_IPostData
}

func (j *jsiiProxy_PostData) DirectWebRemoting() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"directWebRemoting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostData) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostData) GoogleWebToolkit() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"googleWebToolkit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostData) Json() IJsonPostData {
	var returns IJsonPostData
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostData) MultiPartFormData() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multiPartFormData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PostData) Xml() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"xml",
		&returns,
	)
	return returns
}


// Creates an instance of PostData.
//
// Example:
//   const postDataConfig = new PostData({ enabled: true, multiPartFormData: true, xml: true, json: new JsonPostData(), googleWebToolkit: false, directWebRemoting: false });
//   console.log(postDataConfig.enabled); // true
//
func NewPostData(options IPostData) PostData {
	_init_.Initialize()

	j := jsiiProxy_PostData{}

	_jsii_.Create(
		"zap-cdk.PostData",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of PostData.
//
// Example:
//   const postDataConfig = new PostData({ enabled: true, multiPartFormData: true, xml: true, json: new JsonPostData(), googleWebToolkit: false, directWebRemoting: false });
//   console.log(postDataConfig.enabled); // true
//
func NewPostData_Override(p PostData, options IPostData) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PostData",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PostData)SetDirectWebRemoting(val *bool) {
	_jsii_.Set(
		j,
		"directWebRemoting",
		val,
	)
}

func (j *jsiiProxy_PostData)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PostData)SetGoogleWebToolkit(val *bool) {
	_jsii_.Set(
		j,
		"googleWebToolkit",
		val,
	)
}

func (j *jsiiProxy_PostData)SetJson(val IJsonPostData) {
	if err := j.validateSetJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"json",
		val,
	)
}

func (j *jsiiProxy_PostData)SetMultiPartFormData(val *bool) {
	_jsii_.Set(
		j,
		"multiPartFormData",
		val,
	)
}

func (j *jsiiProxy_PostData)SetXml(val *bool) {
	_jsii_.Set(
		j,
		"xml",
		val,
	)
}

