package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for statistics tests.
//
// Example YAML representation:
// ```yaml
// - name: 'test one'                      # Name of the test, optional
//   type: stats                           # Specifies that the test is of type 'stats'
//   statistic: 'stats.addon.something'    # Name of an integer / long statistic
//   site:                                 # Name of the site for site specific tests, supports vars
//   operator: '>='                        # One of '==', '!=', '>=', '>', '<', '<='
//   value: 10                             # Value to compare statistic against
//   onFail: 'info'                        # String: One of 'warn', 'error', 'info', mandatory
// ```.
type IStatisticsTest interface {
	Name() *string
	SetName(n *string)
	OnFail() *string
	SetOnFail(o *string)
	Operator() *string
	SetOperator(o *string)
	Site() *string
	SetSite(s *string)
	Statistic() *string
	SetStatistic(s *string)
	Type() *string
	SetType(t *string)
	Value() *float64
	SetValue(v *float64)
}

// The jsii proxy for IStatisticsTest
type jsiiProxy_IStatisticsTest struct {
	_ byte // padding
}

func (j *jsiiProxy_IStatisticsTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IStatisticsTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_IStatisticsTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_IStatisticsTest) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetSite(val *string) {
	_jsii_.Set(
		j,
		"site",
		val,
	)
}

func (j *jsiiProxy_IStatisticsTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_IStatisticsTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IStatisticsTest) Value() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStatisticsTest)SetValue(val *float64) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

