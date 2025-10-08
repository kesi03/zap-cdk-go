package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPassiveScanParameters interface {
	DisableAllRules() *bool
	SetDisableAllRules(d *bool)
	EnableTags() *bool
	SetEnableTags(e *bool)
	MaxAlertsPerRule() *float64
	SetMaxAlertsPerRule(m *float64)
	MaxBodySizeInBytesToScan() *float64
	SetMaxBodySizeInBytesToScan(m *float64)
	ScanOnlyInScope() *bool
	SetScanOnlyInScope(s *bool)
}

// The jsii proxy for IPassiveScanParameters
type jsiiProxy_IPassiveScanParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IPassiveScanParameters) DisableAllRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disableAllRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanParameters)SetDisableAllRules(val *bool) {
	_jsii_.Set(
		j,
		"disableAllRules",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanParameters) EnableTags() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanParameters)SetEnableTags(val *bool) {
	_jsii_.Set(
		j,
		"enableTags",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanParameters) MaxAlertsPerRule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanParameters)SetMaxAlertsPerRule(val *float64) {
	_jsii_.Set(
		j,
		"maxAlertsPerRule",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanParameters) MaxBodySizeInBytesToScan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBodySizeInBytesToScan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanParameters)SetMaxBodySizeInBytesToScan(val *float64) {
	_jsii_.Set(
		j,
		"maxBodySizeInBytesToScan",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanParameters) ScanOnlyInScope() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanOnlyInScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanParameters)SetScanOnlyInScope(val *bool) {
	_jsii_.Set(
		j,
		"scanOnlyInScope",
		val,
	)
}

