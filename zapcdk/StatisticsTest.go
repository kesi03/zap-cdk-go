package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type StatisticsTest interface {
	IStatisticsTest
	Name() *string
	SetName(val *string)
	OnFail() *string
	SetOnFail(val *string)
	Operator() *string
	SetOperator(val *string)
	Site() *string
	SetSite(val *string)
	Statistic() *string
	SetStatistic(val *string)
	Type() *string
	SetType(val *string)
	Value() *float64
	SetValue(val *float64)
}

// The jsii proxy struct for StatisticsTest
type jsiiProxy_StatisticsTest struct {
	jsiiProxy_IStatisticsTest
}

func (j *jsiiProxy_StatisticsTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatisticsTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatisticsTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatisticsTest) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatisticsTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatisticsTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StatisticsTest) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Creates an instance of StatisticsTest.
func NewStatisticsTest(options IStatisticsTest) StatisticsTest {
	_init_.Initialize()

	if err := validateNewStatisticsTestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_StatisticsTest{}

	_jsii_.Create(
		"zap-cdk.StatisticsTest",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of StatisticsTest.
func NewStatisticsTest_Override(s StatisticsTest, options IStatisticsTest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.StatisticsTest",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_StatisticsTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StatisticsTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_StatisticsTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_StatisticsTest)SetSite(val *string) {
	_jsii_.Set(
		j,
		"site",
		val,
	)
}

func (j *jsiiProxy_StatisticsTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_StatisticsTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_StatisticsTest)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

