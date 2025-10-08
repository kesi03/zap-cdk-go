package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for monitor tests.
//
// Example YAML representation:
// ```yaml
// - name: 'test one'                      # Name of the test, optional
//   type: monitor                         # Specifies that the test is of type 'monitor'
//   statistic: 'stats.addon.something'    # Name of an integer / long statistic
//   site:                                 # Name of the site for site specific tests, supports vars
//   threshold: 10                         # The threshold at which a statistic fails
//   onFail: 'info'                        # String: One of 'warn', 'error', 'info', mandatory
// ```.
type IMonitorTest interface {
	Name() *string
	SetName(n *string)
	OnFail() *string
	SetOnFail(o *string)
	Site() *string
	SetSite(s *string)
	Statistic() *string
	SetStatistic(s *string)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IMonitorTest
type jsiiProxy_IMonitorTest struct {
	_ byte // padding
}

func (j *jsiiProxy_IMonitorTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitorTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IMonitorTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitorTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_IMonitorTest) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitorTest)SetSite(val *string) {
	_jsii_.Set(
		j,
		"site",
		val,
	)
}

func (j *jsiiProxy_IMonitorTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitorTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_IMonitorTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitorTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

