package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IExitStatusParameters interface {
	ErrorExitValue() *float64
	SetErrorExitValue(e *float64)
	ErrorLevel() *string
	SetErrorLevel(e *string)
	OkExitValue() *float64
	SetOkExitValue(o *float64)
	WarnExitValue() *float64
	SetWarnExitValue(w *float64)
	WarnLevel() *string
	SetWarnLevel(w *string)
}

// The jsii proxy for IExitStatusParameters
type jsiiProxy_IExitStatusParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IExitStatusParameters) ErrorExitValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"errorExitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatusParameters)SetErrorExitValue(val *float64) {
	_jsii_.Set(
		j,
		"errorExitValue",
		val,
	)
}

func (j *jsiiProxy_IExitStatusParameters) ErrorLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatusParameters)SetErrorLevel(val *string) {
	_jsii_.Set(
		j,
		"errorLevel",
		val,
	)
}

func (j *jsiiProxy_IExitStatusParameters) OkExitValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"okExitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatusParameters)SetOkExitValue(val *float64) {
	_jsii_.Set(
		j,
		"okExitValue",
		val,
	)
}

func (j *jsiiProxy_IExitStatusParameters) WarnExitValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warnExitValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatusParameters)SetWarnExitValue(val *float64) {
	_jsii_.Set(
		j,
		"warnExitValue",
		val,
	)
}

func (j *jsiiProxy_IExitStatusParameters) WarnLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warnLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatusParameters)SetWarnLevel(val *string) {
	_jsii_.Set(
		j,
		"warnLevel",
		val,
	)
}

