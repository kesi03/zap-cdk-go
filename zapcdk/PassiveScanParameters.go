package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the parameters for configuring a passive scan.
//
// Example:
//   const passiveScanParams = new PassiveScanParameters({
//     maxAlertsPerRule: 5,
//     scanOnlyInScope: true,
//     maxBodySizeInBytesToScan: 0,
//     enableTags: false,
//     disableAllRules: true
//   });
//
type PassiveScanParameters interface {
	IPassiveScanParameters
	DisableAllRules() *bool
	SetDisableAllRules(val *bool)
	EnableTags() *bool
	SetEnableTags(val *bool)
	MaxAlertsPerRule() *float64
	SetMaxAlertsPerRule(val *float64)
	MaxBodySizeInBytesToScan() *float64
	SetMaxBodySizeInBytesToScan(val *float64)
	ScanOnlyInScope() *bool
	SetScanOnlyInScope(val *bool)
}

// The jsii proxy struct for PassiveScanParameters
type jsiiProxy_PassiveScanParameters struct {
	jsiiProxy_IPassiveScanParameters
}

func (j *jsiiProxy_PassiveScanParameters) DisableAllRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"disableAllRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanParameters) EnableTags() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enableTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanParameters) MaxAlertsPerRule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanParameters) MaxBodySizeInBytesToScan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBodySizeInBytesToScan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanParameters) ScanOnlyInScope() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"scanOnlyInScope",
		&returns,
	)
	return returns
}


// Creates an instance of PassiveScanParameters.
func NewPassiveScanParameters(options IPassiveScanParameters) PassiveScanParameters {
	_init_.Initialize()

	if err := validateNewPassiveScanParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PassiveScanParameters{}

	_jsii_.Create(
		"zap-cdk.PassiveScanParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of PassiveScanParameters.
func NewPassiveScanParameters_Override(p PassiveScanParameters, options IPassiveScanParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PassiveScanParameters",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PassiveScanParameters)SetDisableAllRules(val *bool) {
	_jsii_.Set(
		j,
		"disableAllRules",
		val,
	)
}

func (j *jsiiProxy_PassiveScanParameters)SetEnableTags(val *bool) {
	_jsii_.Set(
		j,
		"enableTags",
		val,
	)
}

func (j *jsiiProxy_PassiveScanParameters)SetMaxAlertsPerRule(val *float64) {
	_jsii_.Set(
		j,
		"maxAlertsPerRule",
		val,
	)
}

func (j *jsiiProxy_PassiveScanParameters)SetMaxBodySizeInBytesToScan(val *float64) {
	_jsii_.Set(
		j,
		"maxBodySizeInBytesToScan",
		val,
	)
}

func (j *jsiiProxy_PassiveScanParameters)SetScanOnlyInScope(val *bool) {
	_jsii_.Set(
		j,
		"scanOnlyInScope",
		val,
	)
}

