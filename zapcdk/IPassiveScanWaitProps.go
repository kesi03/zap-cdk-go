package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the PassiveScanWaitConfig construct.
type IPassiveScanWaitProps interface {
	PassiveScanWait() IPassiveScanWait
	SetPassiveScanWait(p IPassiveScanWait)
}

// The jsii proxy for IPassiveScanWaitProps
type jsiiProxy_IPassiveScanWaitProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IPassiveScanWaitProps) PassiveScanWait() IPassiveScanWait {
	var returns IPassiveScanWait
	_jsii_.Get(
		j,
		"passiveScanWait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanWaitProps)SetPassiveScanWait(val IPassiveScanWait) {
	if err := j.validateSetPassiveScanWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveScanWait",
		val,
	)
}

