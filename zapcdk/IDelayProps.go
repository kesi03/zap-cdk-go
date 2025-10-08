package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the DelayConfig construct.
type IDelayProps interface {
	Delay() IDelay
	SetDelay(d IDelay)
}

// The jsii proxy for IDelayProps
type jsiiProxy_IDelayProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IDelayProps) Delay() IDelay {
	var returns IDelay
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelayProps)SetDelay(val IDelay) {
	if err := j.validateSetDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delay",
		val,
	)
}

