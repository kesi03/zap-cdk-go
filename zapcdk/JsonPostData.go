package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for JSON body scanning in POST data.
//
// Example:
//   const jsonConfig = new JsonPostData({ enabled: true, scanNullValues: false });
//   console.log(jsonConfig.enabled); // true
//
type JsonPostData interface {
	IJsonPostData
	// If JSON scanning should be enabled.
	//
	// Default: true.
	Enabled() *bool
	SetEnabled(val *bool)
	// If null values should be scanned.
	//
	// Default: false.
	ScanNullValues() *bool
	SetScanNullValues(val *bool)
}

// The jsii proxy struct for JsonPostData
type jsiiProxy_JsonPostData struct {
	jsiiProxy_IJsonPostData
}

func (j *jsiiProxy_JsonPostData) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JsonPostData) ScanNullValues() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanNullValues",
		&returns,
	)
	return returns
}


// Creates an instance of JsonPostData.
//
// Example:
//   const jsonConfig = new JsonPostData({ enabled: true, scanNullValues: false });
//   console.log(jsonConfig.enabled); // true
//
func NewJsonPostData(options IJsonPostData) JsonPostData {
	_init_.Initialize()

	j := jsiiProxy_JsonPostData{}

	_jsii_.Create(
		"zap-cdk.JsonPostData",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of JsonPostData.
//
// Example:
//   const jsonConfig = new JsonPostData({ enabled: true, scanNullValues: false });
//   console.log(jsonConfig.enabled); // true
//
func NewJsonPostData_Override(j JsonPostData, options IJsonPostData) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.JsonPostData",
		[]interface{}{options},
		j,
	)
}

func (j *jsiiProxy_JsonPostData)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_JsonPostData)SetScanNullValues(val *bool) {
	_jsii_.Set(
		j,
		"scanNullValues",
		val,
	)
}

