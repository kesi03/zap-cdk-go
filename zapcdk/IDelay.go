package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IDelay interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	Parameters() IDelayParameters
	SetParameters(p IDelayParameters)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IDelay
type jsiiProxy_IDelay struct {
	_ byte // padding
}

func (j *jsiiProxy_IDelay) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelay)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_IDelay) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelay)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IDelay) Parameters() IDelayParameters {
	var returns IDelayParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelay)SetParameters(val IDelayParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IDelay) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelay)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

