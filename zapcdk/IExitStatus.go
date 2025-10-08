package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IExitStatus interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	Parameters() IExitStatusParameters
	SetParameters(p IExitStatusParameters)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IExitStatus
type jsiiProxy_IExitStatus struct {
	_ byte // padding
}

func (j *jsiiProxy_IExitStatus) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatus)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_IExitStatus) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatus)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IExitStatus) Parameters() IExitStatusParameters {
	var returns IExitStatusParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatus)SetParameters(val IExitStatusParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IExitStatus) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatus)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

