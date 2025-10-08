package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for alert tests.
//
// Example YAML representation:
// ```yaml
// - name: 'test one'                       # Name of the test, optional
//   type: alert                            # Specifies that the test is of type 'alert'
//   action: passIfPresent                  # String: The condition (presence/absence) of the alert, default: passIfAbsent
//   scanRuleId: 123                        # Integer: The id of the scanRule which generates the alert, mandatory
//   alertName: 'SQL Injection'              # String: The name of the alert generated, optional
//   url: http://www.example.com/path       # String: The url of the request corresponding to the alert generated, optional
//   method: GET                            # String: The method of the request corresponding to the alert generated, optional
//   attack: 'SQL Injection Attack'         # String: The actual attack which generated the alert, optional
//   param: 'username'                      # String: The parameter which was modified to generate the alert, optional
//   evidence: 'Evidence of SQL injection'  # String: The evidence corresponding to the alert generated, optional
//   confidence: High                       # String: The confidence of the alert, one of 'False Positive', 'Low', 'Medium', 'High', 'Confirmed', optional
//   risk: High                             # String: The risk of the alert, one of 'Informational', 'Low', 'Medium', 'High', optional
//   otherInfo: 'Additional context here'   # String: Additional information corresponding to the alert, optional
//   onFail: 'info'                        # String: One of 'warn', 'error', 'info', mandatory
// ```.
type IAlertTest interface {
	Action() *string
	SetAction(a *string)
	AlertName() *string
	SetAlertName(a *string)
	Attack() *string
	SetAttack(a *string)
	Confidence() *string
	SetConfidence(c *string)
	Evidence() *string
	SetEvidence(e *string)
	Method() *string
	SetMethod(m *string)
	Name() *string
	SetName(n *string)
	OnFail() *string
	SetOnFail(o *string)
	Param() *string
	SetParam(p *string)
	Risk() *string
	SetRisk(r *string)
	ScanRuleId() *float64
	SetScanRuleId(s *float64)
	Type() *string
	SetType(t *string)
	Url() *string
	SetUrl(u *string)
}

// The jsii proxy for IAlertTest
type jsiiProxy_IAlertTest struct {
	_ byte // padding
}

func (j *jsiiProxy_IAlertTest) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetAction(val *string) {
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) AlertName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetAlertName(val *string) {
	_jsii_.Set(
		j,
		"alertName",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Attack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetAttack(val *string) {
	_jsii_.Set(
		j,
		"attack",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Confidence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetConfidence(val *string) {
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Evidence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetEvidence(val *string) {
	_jsii_.Set(
		j,
		"evidence",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetMethod(val *string) {
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Param() *string {
	var returns *string
	_jsii_.Get(
		j,
		"param",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetParam(val *string) {
	_jsii_.Set(
		j,
		"param",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Risk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"risk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetRisk(val *string) {
	_jsii_.Set(
		j,
		"risk",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) ScanRuleId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanRuleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetScanRuleId(val *float64) {
	if err := j.validateSetScanRuleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanRuleId",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IAlertTest) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTest)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

