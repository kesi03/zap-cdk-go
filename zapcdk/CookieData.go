package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for cookie data scanning.
//
// Example:
//   const cookieConfig = new CookieData({ enabled: false, encodeCookieValues: false });
//   console.log(cookieConfig.enabled); // false
//
type CookieData interface {
	ICookieData
	// If cookie scanning is enabled.
	//
	// Default: false.
	Enabled() *bool
	SetEnabled(val *bool)
	// If cookie values should be encoded.
	//
	// Default: false.
	EncodeCookieValues() *bool
	SetEncodeCookieValues(val *bool)
}

// The jsii proxy struct for CookieData
type jsiiProxy_CookieData struct {
	jsiiProxy_ICookieData
}

func (j *jsiiProxy_CookieData) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CookieData) EncodeCookieValues() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"encodeCookieValues",
		&returns,
	)
	return returns
}


// Creates an instance of CookieData.
//
// Example:
//   const cookieConfig = new CookieData({ enabled: false, encodeCookieValues: false });
//   console.log(cookieConfig.enabled); // false
//
func NewCookieData(options ICookieData) CookieData {
	_init_.Initialize()

	j := jsiiProxy_CookieData{}

	_jsii_.Create(
		"zap-cdk.CookieData",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of CookieData.
//
// Example:
//   const cookieConfig = new CookieData({ enabled: false, encodeCookieValues: false });
//   console.log(cookieConfig.enabled); // false
//
func NewCookieData_Override(c CookieData, options ICookieData) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.CookieData",
		[]interface{}{options},
		c,
	)
}

func (j *jsiiProxy_CookieData)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CookieData)SetEncodeCookieValues(val *bool) {
	_jsii_.Set(
		j,
		"encodeCookieValues",
		val,
	)
}

