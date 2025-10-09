package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the exit status configuration for the scanning process.
type ExitStatus interface {
	IExitStatus
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Parameters() IExitStatusParameters
	SetParameters(val IExitStatusParameters)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for ExitStatus
type jsiiProxy_ExitStatus struct {
	jsiiProxy_IExitStatus
}

func (j *jsiiProxy_ExitStatus) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatus) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatus) Parameters() IExitStatusParameters {
	var returns IExitStatusParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExitStatus) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of ExitStatus.
func NewExitStatus(options IExitStatus) ExitStatus {
	_init_.Initialize()

	if err := validateNewExitStatusParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExitStatus{}

	_jsii_.Create(
		"zap-cdk.ExitStatus",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ExitStatus.
func NewExitStatus_Override(e ExitStatus, options IExitStatus) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ExitStatus",
		[]interface{}{options},
		e,
	)
}

func (j *jsiiProxy_ExitStatus)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_ExitStatus)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ExitStatus)SetParameters(val IExitStatusParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ExitStatus)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

