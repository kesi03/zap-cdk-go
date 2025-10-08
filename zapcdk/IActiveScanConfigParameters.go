package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScanConfigParameters interface {
	DefaultPolicy() *string
	SetDefaultPolicy(d *string)
	HandleAntiCSRFTokens() *bool
	SetHandleAntiCSRFTokens(h *bool)
	InjectPluginIdInHeader() *bool
	SetInjectPluginIdInHeader(i *bool)
	InputVectors() IInputVectors
	SetInputVectors(i IInputVectors)
	MaxAlertsPerRule() *float64
	SetMaxAlertsPerRule(m *float64)
	MaxRuleDurationInMins() *float64
	SetMaxRuleDurationInMins(m *float64)
	MaxScanDurationInMins() *float64
	SetMaxScanDurationInMins(m *float64)
	ThreadPerHost() *float64
	SetThreadPerHost(t *float64)
}

// The jsii proxy for IActiveScanConfigParameters
type jsiiProxy_IActiveScanConfigParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanConfigParameters) DefaultPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetDefaultPolicy(val *string) {
	_jsii_.Set(
		j,
		"defaultPolicy",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) HandleAntiCSRFTokens() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"handleAntiCSRFTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetHandleAntiCSRFTokens(val *bool) {
	_jsii_.Set(
		j,
		"handleAntiCSRFTokens",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) InjectPluginIdInHeader() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"injectPluginIdInHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetInjectPluginIdInHeader(val *bool) {
	_jsii_.Set(
		j,
		"injectPluginIdInHeader",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) InputVectors() IInputVectors {
	var returns IInputVectors
	_jsii_.Get(
		j,
		"inputVectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetInputVectors(val IInputVectors) {
	if err := j.validateSetInputVectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputVectors",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) MaxAlertsPerRule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetMaxAlertsPerRule(val *float64) {
	_jsii_.Set(
		j,
		"maxAlertsPerRule",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) MaxRuleDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRuleDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetMaxRuleDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxRuleDurationInMins",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) MaxScanDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScanDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetMaxScanDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxScanDurationInMins",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfigParameters) ThreadPerHost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPerHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigParameters)SetThreadPerHost(val *float64) {
	_jsii_.Set(
		j,
		"threadPerHost",
		val,
	)
}

