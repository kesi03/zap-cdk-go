package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the parameters for configuring exit status in the scanning process.
//
// Example:
//   const exitStatusParams = new ExitStatusParameters({
//     errorLevel: 'High',
//     warnLevel: 'Medium',
//     okExitValue: 0,
//     errorExitValue: 1,
//     warnExitValue: 2
//   });
//
type ExitStatusParameters interface {
	IExitStatusParameters
	ErrorExitValue() *float64
	SetErrorExitValue(val *float64)
	ErrorLevel() *string
	SetErrorLevel(val *string)
	OkExitValue() *float64
	SetOkExitValue(val *float64)
	WarnExitValue() *float64
	SetWarnExitValue(val *float64)
	WarnLevel() *string
	SetWarnLevel(val *string)
}

// The jsii proxy struct for ExitStatusParameters
type jsiiProxy_ExitStatusParameters struct {
	jsiiProxy_IExitStatusParameters
}

func (j *jsiiProxy_ExitStatusParameters) ErrorExitValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorExitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatusParameters) ErrorLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatusParameters) OkExitValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"okExitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatusParameters) WarnExitValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warnExitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatusParameters) WarnLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warnLevel",
		&returns,
	)
	return returns
}


// Creates an instance of ExitStatusParameters.
func NewExitStatusParameters(options IExitStatusParameters) ExitStatusParameters {
	_init_.Initialize()

	if err := validateNewExitStatusParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExitStatusParameters{}

	_jsii_.Create(
		"zap-cdk.ExitStatusParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ExitStatusParameters.
func NewExitStatusParameters_Override(e ExitStatusParameters, options IExitStatusParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ExitStatusParameters",
		[]interface{}{options},
		e,
	)
}

func (j *jsiiProxy_ExitStatusParameters)SetErrorExitValue(val *float64) {
	_jsii_.Set(
		j,
		"errorExitValue",
		val,
	)
}

func (j *jsiiProxy_ExitStatusParameters)SetErrorLevel(val *string) {
	_jsii_.Set(
		j,
		"errorLevel",
		val,
	)
}

func (j *jsiiProxy_ExitStatusParameters)SetOkExitValue(val *float64) {
	_jsii_.Set(
		j,
		"okExitValue",
		val,
	)
}

func (j *jsiiProxy_ExitStatusParameters)SetWarnExitValue(val *float64) {
	_jsii_.Set(
		j,
		"warnExitValue",
		val,
	)
}

func (j *jsiiProxy_ExitStatusParameters)SetWarnLevel(val *string) {
	_jsii_.Set(
		j,
		"warnLevel",
		val,
	)
}

