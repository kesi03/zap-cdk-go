package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISessionManagementParametersParameters interface {
	Script() *string
	SetScript(s *string)
	ScriptEngine() *string
	SetScriptEngine(s *string)
}

// The jsii proxy for ISessionManagementParametersParameters
type jsiiProxy_ISessionManagementParametersParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_ISessionManagementParametersParameters) Script() *string {
	var returns *string
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionManagementParametersParameters)SetScript(val *string) {
	_jsii_.Set(
		j,
		"script",
		val,
	)
}

func (j *jsiiProxy_ISessionManagementParametersParameters) ScriptEngine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scriptEngine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionManagementParametersParameters)SetScriptEngine(val *string) {
	_jsii_.Set(
		j,
		"scriptEngine",
		val,
	)
}

