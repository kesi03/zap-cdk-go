package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the parameters for session management in the scanning process.
type SessionManagementParametersParameters interface {
	ISessionManagementParametersParameters
	Script() *string
	SetScript(val *string)
	ScriptEngine() *string
	SetScriptEngine(val *string)
}

// The jsii proxy struct for SessionManagementParametersParameters
type jsiiProxy_SessionManagementParametersParameters struct {
	jsiiProxy_ISessionManagementParametersParameters
}

func (j *jsiiProxy_SessionManagementParametersParameters) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SessionManagementParametersParameters) ScriptEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptEngine",
		&returns,
	)
	return returns
}


// Creates an instance of SessionManagementParametersParameters.
func NewSessionManagementParametersParameters(options ISessionManagementParametersParameters) SessionManagementParametersParameters {
	_init_.Initialize()

	j := jsiiProxy_SessionManagementParametersParameters{}

	_jsii_.Create(
		"zap-cdk.SessionManagementParametersParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of SessionManagementParametersParameters.
func NewSessionManagementParametersParameters_Override(s SessionManagementParametersParameters, options ISessionManagementParametersParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SessionManagementParametersParameters",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SessionManagementParametersParameters)SetScript(val *string) {
	_jsii_.Set(
		j,
		"script",
		val,
	)
}

func (j *jsiiProxy_SessionManagementParametersParameters)SetScriptEngine(val *string) {
	_jsii_.Set(
		j,
		"scriptEngine",
		val,
	)
}

