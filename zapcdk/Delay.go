package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type Delay interface {
	IDelay
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Parameters() IDelayParameters
	SetParameters(val IDelayParameters)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for Delay
type jsiiProxy_Delay struct {
	jsiiProxy_IDelay
}

func (j *jsiiProxy_Delay) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Delay) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Delay) Parameters() IDelayParameters {
	var returns IDelayParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Delay) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of Delay.
func NewDelay(parameters IDelayParameters, enabled *bool, alwaysRun *bool) Delay {
	_init_.Initialize()

	if err := validateNewDelayParameters(parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_Delay{}

	_jsii_.Create(
		"zap-cdk.Delay",
		[]interface{}{parameters, enabled, alwaysRun},
		&j,
	)

	return &j
}

// Creates an instance of Delay.
func NewDelay_Override(d Delay, parameters IDelayParameters, enabled *bool, alwaysRun *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Delay",
		[]interface{}{parameters, enabled, alwaysRun},
		d,
	)
}

func (j *jsiiProxy_Delay)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_Delay)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_Delay)SetParameters(val IDelayParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_Delay)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

