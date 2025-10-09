package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type DelayParameters interface {
	IDelayParameters
	FileName() *string
	SetFileName(val *string)
	Time() *string
	SetTime(val *string)
}

// The jsii proxy struct for DelayParameters
type jsiiProxy_DelayParameters struct {
	jsiiProxy_IDelayParameters
}

func (j *jsiiProxy_DelayParameters) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DelayParameters) Time() *string {
	var returns *string
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}


// Creates an instance of DelayParameters.
func NewDelayParameters(time *string, fileName *string) DelayParameters {
	_init_.Initialize()

	j := jsiiProxy_DelayParameters{}

	_jsii_.Create(
		"zap-cdk.DelayParameters",
		[]interface{}{time, fileName},
		&j,
	)

	return &j
}

// Creates an instance of DelayParameters.
func NewDelayParameters_Override(d DelayParameters, time *string, fileName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.DelayParameters",
		[]interface{}{time, fileName},
		d,
	)
}

func (j *jsiiProxy_DelayParameters)SetFileName(val *string) {
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_DelayParameters)SetTime(val *string) {
	_jsii_.Set(
		j,
		"time",
		val,
	)
}

