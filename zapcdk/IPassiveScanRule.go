package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPassiveScanRule interface {
	Id() *float64
	SetId(i *float64)
	Name() *string
	SetName(n *string)
	Threshold() *string
	SetThreshold(t *string)
}

// The jsii proxy for IPassiveScanRule
type jsiiProxy_IPassiveScanRule struct {
	_ byte // padding
}

func (j *jsiiProxy_IPassiveScanRule) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanRule)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanRule)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanRule) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanRule)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

