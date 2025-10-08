package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPassiveScanWait interface {
	MaxDuration() *float64
	SetMaxDuration(m *float64)
}

// The jsii proxy for IPassiveScanWait
type jsiiProxy_IPassiveScanWait struct {
	_ byte // padding
}

func (j *jsiiProxy_IPassiveScanWait) MaxDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanWait)SetMaxDuration(val *float64) {
	_jsii_.Set(
		j,
		"maxDuration",
		val,
	)
}

