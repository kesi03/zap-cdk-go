package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAlertFilter interface {
	Attack() *string
	SetAttack(a *string)
	AttackRegex() *bool
	SetAttackRegex(a *bool)
	Context() *string
	SetContext(c *string)
	Evidence() *string
	SetEvidence(e *string)
	EvidenceRegex() *bool
	SetEvidenceRegex(e *bool)
	NewRisk() *string
	SetNewRisk(n *string)
	Parameter() *string
	SetParameter(p *string)
	ParameterRegex() *bool
	SetParameterRegex(p *bool)
	RuleId() *float64
	SetRuleId(r *float64)
	Url() *string
	SetUrl(u *string)
	UrlRegex() *bool
	SetUrlRegex(u *bool)
}

// The jsii proxy for IAlertFilter
type jsiiProxy_IAlertFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_IAlertFilter) Attack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetAttack(val *string) {
	_jsii_.Set(
		j,
		"attack",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) AttackRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"attackRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetAttackRegex(val *bool) {
	_jsii_.Set(
		j,
		"attackRegex",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) Evidence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetEvidence(val *string) {
	_jsii_.Set(
		j,
		"evidence",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) EvidenceRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"evidenceRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetEvidenceRegex(val *bool) {
	_jsii_.Set(
		j,
		"evidenceRegex",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) NewRisk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newRisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetNewRisk(val *string) {
	if err := j.validateSetNewRiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newRisk",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) Parameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetParameter(val *string) {
	_jsii_.Set(
		j,
		"parameter",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) ParameterRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"parameterRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetParameterRegex(val *bool) {
	_jsii_.Set(
		j,
		"parameterRegex",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) RuleId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ruleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetRuleId(val *float64) {
	if err := j.validateSetRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleId",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_IAlertFilter) UrlRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"urlRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertFilter)SetUrlRegex(val *bool) {
	_jsii_.Set(
		j,
		"urlRegex",
		val,
	)
}

