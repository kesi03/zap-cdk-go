package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the parameters for authentication in the scanning process.
type AuthenticationParametersParameters interface {
	IAuthenticationParametersParameters
	Hostname() *string
	SetHostname(val *string)
	LoginPageUrl() *string
	SetLoginPageUrl(val *string)
	LoginRequestBody() *string
	SetLoginRequestBody(val *string)
	LoginRequestUrl() *string
	SetLoginRequestUrl(val *string)
	Port() *float64
	SetPort(val *float64)
	Realm() *string
	SetRealm(val *string)
	Script() *string
	SetScript(val *string)
	ScriptEngine() *string
	SetScriptEngine(val *string)
	ScriptInline() *string
	SetScriptInline(val *string)
}

// The jsii proxy struct for AuthenticationParametersParameters
type jsiiProxy_AuthenticationParametersParameters struct {
	jsiiProxy_IAuthenticationParametersParameters
}

func (j *jsiiProxy_AuthenticationParametersParameters) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) LoginPageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) LoginRequestBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginRequestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) LoginRequestUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginRequestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) ScriptEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersParameters) ScriptInline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptInline",
		&returns,
	)
	return returns
}


// Creates an instance of AuthenticationParametersParameters.
func NewAuthenticationParametersParameters(options IAuthenticationParametersParameters) AuthenticationParametersParameters {
	_init_.Initialize()

	j := jsiiProxy_AuthenticationParametersParameters{}

	_jsii_.Create(
		"zap-cdk.AuthenticationParametersParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AuthenticationParametersParameters.
func NewAuthenticationParametersParameters_Override(a AuthenticationParametersParameters, options IAuthenticationParametersParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AuthenticationParametersParameters",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetLoginPageUrl(val *string) {
	_jsii_.Set(
		j,
		"loginPageUrl",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetLoginRequestBody(val *string) {
	_jsii_.Set(
		j,
		"loginRequestBody",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetLoginRequestUrl(val *string) {
	_jsii_.Set(
		j,
		"loginRequestUrl",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetRealm(val *string) {
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetScript(val *string) {
	_jsii_.Set(
		j,
		"script",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetScriptEngine(val *string) {
	_jsii_.Set(
		j,
		"scriptEngine",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersParameters)SetScriptInline(val *string) {
	_jsii_.Set(
		j,
		"scriptInline",
		val,
	)
}

