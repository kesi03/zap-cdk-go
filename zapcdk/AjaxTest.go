package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type AjaxTest interface {
	IAjaxTest
	Name() *string
	SetName(val *string)
	OnFail() *string
	SetOnFail(val *string)
	Operator() *string
	SetOperator(val *string)
	Statistic() *string
	SetStatistic(val *string)
	Type() *string
	SetType(val *string)
	Value() *float64
	SetValue(val *float64)
}

// The jsii proxy struct for AjaxTest
type jsiiProxy_AjaxTest struct {
	jsiiProxy_IAjaxTest
}

func (j *jsiiProxy_AjaxTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AjaxTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AjaxTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AjaxTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AjaxTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AjaxTest) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates an instance of AjaxTest.
//
// Example:
//   const ajaxTest = new AjaxTest({
//     name: 'Check AJAX requests',
//     type: 'stats',
//     statistic: 'ajax.requests',
//     operator: '>',
//     value: 10,
//     onFail: 'warn'
//   });
//
func NewAjaxTest(options IAjaxTest) AjaxTest {
	_init_.Initialize()

	if err := validateNewAjaxTestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AjaxTest{}

	_jsii_.Create(
		"zap-cdk.AjaxTest",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AjaxTest.
//
// Example:
//   const ajaxTest = new AjaxTest({
//     name: 'Check AJAX requests',
//     type: 'stats',
//     statistic: 'ajax.requests',
//     operator: '>',
//     value: 10,
//     onFail: 'warn'
//   });
//
func NewAjaxTest_Override(a AjaxTest, options IAjaxTest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AjaxTest",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AjaxTest)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AjaxTest)SetOnFail(val *string) {
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_AjaxTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_AjaxTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_AjaxTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AjaxTest)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

