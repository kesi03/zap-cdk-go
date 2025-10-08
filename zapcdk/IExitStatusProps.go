package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ExitStatusConfig construct.
type IExitStatusProps interface {
	ExitStatus() IExitStatus
	SetExitStatus(e IExitStatus)
}

// The jsii proxy for IExitStatusProps
type jsiiProxy_IExitStatusProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IExitStatusProps) ExitStatus() IExitStatus {
	var returns IExitStatus
	_jsii_.Get(
		j,
		"exitStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExitStatusProps)SetExitStatus(val IExitStatus) {
	if err := j.validateSetExitStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exitStatus",
		val,
	)
}

