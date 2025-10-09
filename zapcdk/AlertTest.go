package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type AlertTest interface {
	IAlertTest
	Action() *string
	SetAction(val *string)
	AlertName() *string
	SetAlertName(val *string)
	Attack() *string
	SetAttack(val *string)
	Confidence() *string
	SetConfidence(val *string)
	Evidence() *string
	SetEvidence(val *string)
	Method() *string
	SetMethod(val *string)
	Name() *string
	SetName(val *string)
	OnFail() *string
	SetOnFail(val *string)
	Param() *string
	SetParam(val *string)
	Risk() *string
	SetRisk(val *string)
	ScanRuleId() *float64
	SetScanRuleId(val *float64)
	Type() *string
	SetType(val *string)
	Url() *string
	SetUrl(val *string)
}

// The jsii proxy struct for AlertTest
type jsiiProxy_AlertTest struct {
	jsiiProxy_IAlertTest
}

func (j *jsiiProxy_AlertTest) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) AlertName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Attack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Confidence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Evidence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Param() *string {
	var returns *string
	_jsii_.Get(
		j,
		"param",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Risk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"risk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) ScanRuleId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTest) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Creates an instance of AlertTest.
//
// Example:
//   const alertTest = new AlertTest({
//     name: 'test one',
//     action: 'passIfPresent',
//     scanRuleId: 123,
//     alertName: 'SQL Injection',
//     url: 'http://www.example.com/path',
//     method: 'GET',
//     attack: 'SQL Injection Attack',
//     param: 'username',
//     evidence: 'Evidence of SQL injection',
//     confidence: 'High',
//     risk: 'High',
//     onFail: 'info'
//   });
//
func NewAlertTest(options IAlertTest) AlertTest {
	_init_.Initialize()

	if err := validateNewAlertTestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertTest{}

	_jsii_.Create(
		"zap-cdk.AlertTest",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AlertTest.
//
// Example:
//   const alertTest = new AlertTest({
//     name: 'test one',
//     action: 'passIfPresent',
//     scanRuleId: 123,
//     alertName: 'SQL Injection',
//     url: 'http://www.example.com/path',
//     method: 'GET',
//     attack: 'SQL Injection Attack',
//     param: 'username',
//     evidence: 'Evidence of SQL injection',
//     confidence: 'High',
//     risk: 'High',
//     onFail: 'info'
//   });
//
func NewAlertTest_Override(a AlertTest, options IAlertTest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AlertTest",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AlertTest)SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetAlertName(val *string) {
	_jsii_.Set(
		j,
		"alertName",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetAttack(val *string) {
	_jsii_.Set(
		j,
		"attack",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetConfidence(val *string) {
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetEvidence(val *string) {
	_jsii_.Set(
		j,
		"evidence",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetMethod(val *string) {
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetParam(val *string) {
	_jsii_.Set(
		j,
		"param",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetRisk(val *string) {
	_jsii_.Set(
		j,
		"risk",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetScanRuleId(val *float64) {
	if err := j.validateSetScanRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanRuleId",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AlertTest)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

