package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a test configuration.
type IAjaxTest interface {
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

// The jsii proxy for IAjaxTest
type jsiiProxy_IAjaxTest struct {
	_ byte // padding
}

func (j *jsiiProxy_IAjaxTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAjaxTest)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IAjaxTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAjaxTest)SetOnFail(val *string) {
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_IAjaxTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAjaxTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_IAjaxTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAjaxTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_IAjaxTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAjaxTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IAjaxTest) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAjaxTest)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

