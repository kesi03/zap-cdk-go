package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a test configuration for the spider.
type ISpiderTest interface {
	Name() *string
	SetName(n *string)
	OnFail() *string
	SetOnFail(o *string)
	Operator() *string
	SetOperator(o *string)
	Statistic() *string
	SetStatistic(s *string)
	Type() *string
	SetType(t *string)
	Value() *float64
	SetValue(v *float64)
}

// The jsii proxy for ISpiderTest
type jsiiProxy_ISpiderTest struct {
	_ byte // padding
}

func (j *jsiiProxy_ISpiderTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ISpiderTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_ISpiderTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_ISpiderTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_ISpiderTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ISpiderTest) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpiderTest)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

