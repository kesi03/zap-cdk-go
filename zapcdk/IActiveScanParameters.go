package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScanParameters interface {
	AddQueryParam() *bool
	SetAddQueryParam(a *bool)
	Context() *string
	SetContext(c *string)
	DefaultPolicy() *string
	SetDefaultPolicy(d *string)
	DelayInMs() *float64
	SetDelayInMs(d *float64)
	HandleAntiCSRFTokens() *bool
	SetHandleAntiCSRFTokens(h *bool)
	InjectPluginIdInHeader() *bool
	SetInjectPluginIdInHeader(i *bool)
	MaxAlertsPerRule() *float64
	SetMaxAlertsPerRule(m *float64)
	MaxRuleDurationInMins() *float64
	SetMaxRuleDurationInMins(m *float64)
	MaxScanDurationInMins() *float64
	SetMaxScanDurationInMins(m *float64)
	Policy() *string
	SetPolicy(p *string)
	ScanHeadersAllRequests() *bool
	SetScanHeadersAllRequests(s *bool)
	Tests() *[]interface{}
	SetTests(t *[]interface{})
	ThreadPerHost() *float64
	SetThreadPerHost(t *float64)
	Url() *string
	SetUrl(u *string)
	User() *string
	SetUser(u *string)
}

// The jsii proxy for IActiveScanParameters
type jsiiProxy_IActiveScanParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanParameters) AddQueryParam() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addQueryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetAddQueryParam(val *bool) {
	_jsii_.Set(
		j,
		"addQueryParam",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) DefaultPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetDefaultPolicy(val *string) {
	_jsii_.Set(
		j,
		"defaultPolicy",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) DelayInMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayInMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetDelayInMs(val *float64) {
	_jsii_.Set(
		j,
		"delayInMs",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) HandleAntiCSRFTokens() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"handleAntiCSRFTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetHandleAntiCSRFTokens(val *bool) {
	_jsii_.Set(
		j,
		"handleAntiCSRFTokens",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) InjectPluginIdInHeader() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"injectPluginIdInHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetInjectPluginIdInHeader(val *bool) {
	_jsii_.Set(
		j,
		"injectPluginIdInHeader",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) MaxAlertsPerRule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetMaxAlertsPerRule(val *float64) {
	_jsii_.Set(
		j,
		"maxAlertsPerRule",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) MaxRuleDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRuleDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetMaxRuleDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxRuleDurationInMins",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) MaxScanDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScanDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetMaxScanDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxScanDurationInMins",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) ScanHeadersAllRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanHeadersAllRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetScanHeadersAllRequests(val *bool) {
	_jsii_.Set(
		j,
		"scanHeadersAllRequests",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) Tests() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"tests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetTests(val *[]interface{}) {
	if err := j.validateSetTestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tests",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) ThreadPerHost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPerHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetThreadPerHost(val *float64) {
	_jsii_.Set(
		j,
		"threadPerHost",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_IActiveScanParameters) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanParameters)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

