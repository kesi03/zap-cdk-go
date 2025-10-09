package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing parameters for an active scan.
type ActiveScanParameters interface {
	IActiveScanParameters
	AddQueryParam() *bool
	SetAddQueryParam(val *bool)
	Context() *string
	SetContext(val *string)
	DefaultPolicy() *string
	SetDefaultPolicy(val *string)
	DelayInMs() *float64
	SetDelayInMs(val *float64)
	HandleAntiCSRFTokens() *bool
	SetHandleAntiCSRFTokens(val *bool)
	InjectPluginIdInHeader() *bool
	SetInjectPluginIdInHeader(val *bool)
	MaxAlertsPerRule() *float64
	SetMaxAlertsPerRule(val *float64)
	MaxRuleDurationInMins() *float64
	SetMaxRuleDurationInMins(val *float64)
	MaxScanDurationInMins() *float64
	SetMaxScanDurationInMins(val *float64)
	Policy() *string
	SetPolicy(val *string)
	ScanHeadersAllRequests() *bool
	SetScanHeadersAllRequests(val *bool)
	Tests() *[]interface{}
	SetTests(val *[]interface{})
	ThreadPerHost() *float64
	SetThreadPerHost(val *float64)
	Url() *string
	SetUrl(val *string)
	User() *string
	SetUser(val *string)
}

// The jsii proxy struct for ActiveScanParameters
type jsiiProxy_ActiveScanParameters struct {
	jsiiProxy_IActiveScanParameters
}

func (j *jsiiProxy_ActiveScanParameters) AddQueryParam() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addQueryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) DefaultPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) DelayInMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delayInMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) HandleAntiCSRFTokens() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"handleAntiCSRFTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) InjectPluginIdInHeader() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"injectPluginIdInHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) MaxAlertsPerRule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) MaxRuleDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRuleDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) MaxScanDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScanDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) ScanHeadersAllRequests() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanHeadersAllRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) Tests() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"tests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) ThreadPerHost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPerHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanParameters) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanParameters.
func NewActiveScanParameters(options IActiveScanParameters) ActiveScanParameters {
	_init_.Initialize()

	j := jsiiProxy_ActiveScanParameters{}

	_jsii_.Create(
		"zap-cdk.ActiveScanParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanParameters.
func NewActiveScanParameters_Override(a ActiveScanParameters, options IActiveScanParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanParameters",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetAddQueryParam(val *bool) {
	_jsii_.Set(
		j,
		"addQueryParam",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetDefaultPolicy(val *string) {
	_jsii_.Set(
		j,
		"defaultPolicy",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetDelayInMs(val *float64) {
	_jsii_.Set(
		j,
		"delayInMs",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetHandleAntiCSRFTokens(val *bool) {
	_jsii_.Set(
		j,
		"handleAntiCSRFTokens",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetInjectPluginIdInHeader(val *bool) {
	_jsii_.Set(
		j,
		"injectPluginIdInHeader",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetMaxAlertsPerRule(val *float64) {
	_jsii_.Set(
		j,
		"maxAlertsPerRule",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetMaxRuleDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxRuleDurationInMins",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetMaxScanDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxScanDurationInMins",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetPolicy(val *string) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetScanHeadersAllRequests(val *bool) {
	_jsii_.Set(
		j,
		"scanHeadersAllRequests",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetTests(val *[]interface{}) {
	if err := j.validateSetTestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tests",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetThreadPerHost(val *float64) {
	_jsii_.Set(
		j,
		"threadPerHost",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_ActiveScanParameters)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

