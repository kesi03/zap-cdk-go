package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a test configuration for the spider.
//
// Example:
//   const spiderTest = new SpiderTest({
//     statistic: 'urls',
//     operator: '>=',
//     value: 10,
//     onFail: 'error'
//   });
//
type SpiderTest interface {
	ISpiderTest
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

// The jsii proxy struct for SpiderTest
type jsiiProxy_SpiderTest struct {
	jsiiProxy_ISpiderTest
}

func (j *jsiiProxy_SpiderTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpiderTest) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates an instance of SpiderTest.
//
// Example:
//   const spiderTest = new SpiderTest({
//     statistic: 'urls',
//     operator: '>=',
//     value: 10,
//     onFail: 'error'
//   });
//
func NewSpiderTest(options ISpiderTest) SpiderTest {
	_init_.Initialize()

	if err := validateNewSpiderTestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpiderTest{}

	_jsii_.Create(
		"zap-cdk.SpiderTest",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of SpiderTest.
//
// Example:
//   const spiderTest = new SpiderTest({
//     statistic: 'urls',
//     operator: '>=',
//     value: 10,
//     onFail: 'error'
//   });
//
func NewSpiderTest_Override(s SpiderTest, options ISpiderTest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SpiderTest",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SpiderTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SpiderTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_SpiderTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_SpiderTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_SpiderTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_SpiderTest)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

