package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IRule interface {
	Id() *float64
	SetId(i *float64)
	Name() *string
	SetName(n *string)
	Strength() *string
	SetStrength(s *string)
	Threshold() *string
	SetThreshold(t *string)
}

// The jsii proxy for IRule
type jsiiProxy_IRule struct {
	_ byte // padding
}

func (j *jsiiProxy_IRule) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IRule) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_IRule) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRule)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

