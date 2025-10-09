package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a filter for alerts in the scanning process.
//
// Example:
//   const alertFilter = new AlertFilter({
//     ruleId: 10010,
//     newRisk: 'Low',
//     context: 'MyContext',
//     url: '.*example.*',
//     urlRegex: true,
//     parameter: 'sessionid',
//     parameterRegex: false,
//     attack: 'SQL Injection',
//     attackRegex: false,
//     evidence: 'SELECT',
//     evidenceRegex: true
//   });
//
type AlertFilter interface {
	IAlertFilter
	Attack() *string
	SetAttack(val *string)
	AttackRegex() *bool
	SetAttackRegex(val *bool)
	Context() *string
	SetContext(val *string)
	Evidence() *string
	SetEvidence(val *string)
	EvidenceRegex() *bool
	SetEvidenceRegex(val *bool)
	NewRisk() *string
	SetNewRisk(val *string)
	Parameter() *string
	SetParameter(val *string)
	ParameterRegex() *bool
	SetParameterRegex(val *bool)
	RuleId() *float64
	SetRuleId(val *float64)
	Url() *string
	SetUrl(val *string)
	UrlRegex() *bool
	SetUrlRegex(val *bool)
}

// The jsii proxy struct for AlertFilter
type jsiiProxy_AlertFilter struct {
	jsiiProxy_IAlertFilter
}

func (j *jsiiProxy_AlertFilter) Attack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) AttackRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"attackRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) Evidence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) EvidenceRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"evidenceRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) NewRisk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newRisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) Parameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) ParameterRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parameterRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) RuleId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertFilter) UrlRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"urlRegex",
		&returns,
	)
	return returns
}


// Creates an instance of AlertFilter.
//
// Example:
//   const alertFilter = new AlertFilter({
//     ruleId: 10010,
//     newRisk: 'Low',
//     context: 'MyContext',
//     url: '.*example.*',
//     urlRegex: true,
//     parameter: 'sessionid',
//     parameterRegex: false,
//     attack: 'SQL Injection',
//     attackRegex: false,
//     evidence: 'SELECT',
//     evidenceRegex: true
//   });
//
func NewAlertFilter(options IAlertFilter) AlertFilter {
	_init_.Initialize()

	if err := validateNewAlertFilterParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertFilter{}

	_jsii_.Create(
		"zap-cdk.AlertFilter",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AlertFilter.
//
// Example:
//   const alertFilter = new AlertFilter({
//     ruleId: 10010,
//     newRisk: 'Low',
//     context: 'MyContext',
//     url: '.*example.*',
//     urlRegex: true,
//     parameter: 'sessionid',
//     parameterRegex: false,
//     attack: 'SQL Injection',
//     attackRegex: false,
//     evidence: 'SELECT',
//     evidenceRegex: true
//   });
//
func NewAlertFilter_Override(a AlertFilter, options IAlertFilter) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AlertFilter",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AlertFilter)SetAttack(val *string) {
	_jsii_.Set(
		j,
		"attack",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetAttackRegex(val *bool) {
	_jsii_.Set(
		j,
		"attackRegex",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetEvidence(val *string) {
	_jsii_.Set(
		j,
		"evidence",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetEvidenceRegex(val *bool) {
	_jsii_.Set(
		j,
		"evidenceRegex",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetNewRisk(val *string) {
	if err := j.validateSetNewRiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newRisk",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetParameter(val *string) {
	_jsii_.Set(
		j,
		"parameter",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetParameterRegex(val *bool) {
	_jsii_.Set(
		j,
		"parameterRegex",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetRuleId(val *float64) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_AlertFilter)SetUrlRegex(val *bool) {
	_jsii_.Set(
		j,
		"urlRegex",
		val,
	)
}

