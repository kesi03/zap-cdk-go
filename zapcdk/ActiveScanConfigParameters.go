package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the parameters for configuring an active scan.
//
// Example:
//   const activeScanConfig = new ActiveScanConfigParameters({
//     maxRuleDurationInMins: 0,
//     maxScanDurationInMins: 0,
//     maxAlertsPerRule: 0,
//     defaultPolicy: 'Default Policy',
//     handleAntiCSRFTokens: false,
//     injectPluginIdInHeader: false,
//     threadPerHost: 4,
//     inputVectors: new InputVectors()
//   });
//   console.log(activeScanConfig.defaultPolicy); // 'Default Policy'
//
type ActiveScanConfigParameters interface {
	IActiveScanConfigParameters
	DefaultPolicy() *string
	SetDefaultPolicy(val *string)
	HandleAntiCSRFTokens() *bool
	SetHandleAntiCSRFTokens(val *bool)
	InjectPluginIdInHeader() *bool
	SetInjectPluginIdInHeader(val *bool)
	InputVectors() IInputVectors
	SetInputVectors(val IInputVectors)
	MaxAlertsPerRule() *float64
	SetMaxAlertsPerRule(val *float64)
	MaxRuleDurationInMins() *float64
	SetMaxRuleDurationInMins(val *float64)
	MaxScanDurationInMins() *float64
	SetMaxScanDurationInMins(val *float64)
	ThreadPerHost() *float64
	SetThreadPerHost(val *float64)
}

// The jsii proxy struct for ActiveScanConfigParameters
type jsiiProxy_ActiveScanConfigParameters struct {
	jsiiProxy_IActiveScanConfigParameters
}

func (j *jsiiProxy_ActiveScanConfigParameters) DefaultPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) HandleAntiCSRFTokens() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"handleAntiCSRFTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) InjectPluginIdInHeader() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"injectPluginIdInHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) InputVectors() IInputVectors {
	var returns IInputVectors
	_jsii_.Get(
		j,
		"inputVectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) MaxAlertsPerRule() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAlertsPerRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) MaxRuleDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRuleDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) MaxScanDurationInMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScanDurationInMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfigParameters) ThreadPerHost() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadPerHost",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanConfigParameters.
//
// Example:
//   const activeScanConfig = new ActiveScanConfigParameters({
//     maxRuleDurationInMins: 0,
//     maxScanDurationInMins: 0,
//     maxAlertsPerRule: 0,
//     defaultPolicy: 'Default Policy',
//     handleAntiCSRFTokens: false,
//     injectPluginIdInHeader: false,
//     threadPerHost: 4,
//     inputVectors: new InputVectors()
//   });
//   console.log(activeScanConfig.defaultPolicy); // 'Default Policy'
//
func NewActiveScanConfigParameters(options IActiveScanConfigParameters) ActiveScanConfigParameters {
	_init_.Initialize()

	j := jsiiProxy_ActiveScanConfigParameters{}

	_jsii_.Create(
		"zap-cdk.ActiveScanConfigParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanConfigParameters.
//
// Example:
//   const activeScanConfig = new ActiveScanConfigParameters({
//     maxRuleDurationInMins: 0,
//     maxScanDurationInMins: 0,
//     maxAlertsPerRule: 0,
//     defaultPolicy: 'Default Policy',
//     handleAntiCSRFTokens: false,
//     injectPluginIdInHeader: false,
//     threadPerHost: 4,
//     inputVectors: new InputVectors()
//   });
//   console.log(activeScanConfig.defaultPolicy); // 'Default Policy'
//
func NewActiveScanConfigParameters_Override(a ActiveScanConfigParameters, options IActiveScanConfigParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanConfigParameters",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetDefaultPolicy(val *string) {
	_jsii_.Set(
		j,
		"defaultPolicy",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetHandleAntiCSRFTokens(val *bool) {
	_jsii_.Set(
		j,
		"handleAntiCSRFTokens",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetInjectPluginIdInHeader(val *bool) {
	_jsii_.Set(
		j,
		"injectPluginIdInHeader",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetInputVectors(val IInputVectors) {
	if err := j.validateSetInputVectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputVectors",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetMaxAlertsPerRule(val *float64) {
	_jsii_.Set(
		j,
		"maxAlertsPerRule",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetMaxRuleDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxRuleDurationInMins",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetMaxScanDurationInMins(val *float64) {
	_jsii_.Set(
		j,
		"maxScanDurationInMins",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfigParameters)SetThreadPerHost(val *float64) {
	_jsii_.Set(
		j,
		"threadPerHost",
		val,
	)
}

