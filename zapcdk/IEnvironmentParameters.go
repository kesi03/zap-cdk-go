package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IEnvironmentParameters interface {
	ContinueOnFailure() *bool
	SetContinueOnFailure(c *bool)
	FailOnError() *bool
	SetFailOnError(f *bool)
	FailOnWarning() *bool
	SetFailOnWarning(f *bool)
	ProgressToStdout() *bool
	SetProgressToStdout(p *bool)
}

// The jsii proxy for IEnvironmentParameters
type jsiiProxy_IEnvironmentParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IEnvironmentParameters) ContinueOnFailure() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"continueOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentParameters)SetContinueOnFailure(val *bool) {
	_jsii_.Set(
		j,
		"continueOnFailure",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentParameters) FailOnError() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"failOnError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentParameters)SetFailOnError(val *bool) {
	_jsii_.Set(
		j,
		"failOnError",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentParameters) FailOnWarning() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"failOnWarning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentParameters)SetFailOnWarning(val *bool) {
	_jsii_.Set(
		j,
		"failOnWarning",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentParameters) ProgressToStdout() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"progressToStdout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentParameters)SetProgressToStdout(val *bool) {
	_jsii_.Set(
		j,
		"progressToStdout",
		val,
	)
}

