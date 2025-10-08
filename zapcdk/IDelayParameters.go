package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IDelayParameters interface {
	FileName() *string
	SetFileName(f *string)
	Time() *string
	SetTime(t *string)
}

// The jsii proxy for IDelayParameters
type jsiiProxy_IDelayParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IDelayParameters) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelayParameters)SetFileName(val *string) {
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_IDelayParameters) Time() *string {
	var returns *string
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDelayParameters)SetTime(val *string) {
	_jsii_.Set(
		j,
		"time",
		val,
	)
}

