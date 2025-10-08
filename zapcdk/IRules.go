package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IRules interface {
	Id() *float64
	SetId(i *float64)
	Name() *string
	SetName(n *string)
	Strength() *string
	SetStrength(s *string)
	Threshold() *string
	SetThreshold(t *string)
}

// The jsii proxy for IRules
type jsiiProxy_IRules struct {
	_ byte // padding
}

func (j *jsiiProxy_IRules) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IRules) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IRules) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_IRules) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRules)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

