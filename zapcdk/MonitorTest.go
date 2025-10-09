package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a monitor test.
//
// Example:
//   const monitorTest = new MonitorTest({
//     name: 'test one',
//     statistic: 'stats.addon.something',
//     site: 'MySite',
//     onFail: 'info'
//   });
//
type MonitorTest interface {
	IMonitorTest
	Name() *string
	SetName(val *string)
	OnFail() *string
	SetOnFail(val *string)
	Site() *string
	SetSite(val *string)
	Statistic() *string
	SetStatistic(val *string)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for MonitorTest
type jsiiProxy_MonitorTest struct {
	jsiiProxy_IMonitorTest
}

func (j *jsiiProxy_MonitorTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorTest) Site() *string {
	var returns *string
	_jsii_.Get(
		j,
		"site",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorTest) Statistic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statistic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of MonitorTest.
func NewMonitorTest(options IMonitorTest) MonitorTest {
	_init_.Initialize()

	if err := validateNewMonitorTestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorTest{}

	_jsii_.Create(
		"zap-cdk.MonitorTest",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of MonitorTest.
func NewMonitorTest_Override(m MonitorTest, options IMonitorTest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.MonitorTest",
		[]interface{}{options},
		m,
	)
}

func (j *jsiiProxy_MonitorTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_MonitorTest)SetSite(val *string) {
	_jsii_.Set(
		j,
		"site",
		val,
	)
}

func (j *jsiiProxy_MonitorTest)SetStatistic(val *string) {
	if err := j.validateSetStatisticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statistic",
		val,
	)
}

func (j *jsiiProxy_MonitorTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

