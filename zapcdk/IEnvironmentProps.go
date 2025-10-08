package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the EnvironmentConfig construct.
type IEnvironmentProps interface {
	Environment() IEnvironment
	SetEnvironment(e IEnvironment)
}

// The jsii proxy for IEnvironmentProps
type jsiiProxy_IEnvironmentProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IEnvironmentProps) Environment() IEnvironment {
	var returns IEnvironment
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProps)SetEnvironment(val IEnvironment) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

