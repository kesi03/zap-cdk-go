package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the passive scan configuration.
//
// Example:
//   const passiveScanConfig = new PassiveScanConfig({
//     parameters: {
//       maxAlertsPerRule: 5,
//       scanOnlyInScope: true,
//       maxBodySizeInBytesToScan: 0,
//       enableTags: false,
//       disableAllRules: true
//     },
//     rules: [
//       { id: 10010, name: 'Cross-Domain Misconfiguration', threshold: 'Medium' },
//       { id: 10011, name: 'CSP Header Not Set', threshold: 'High' }
//     ],
//     enabled: true,
//     alwaysRun: false
//   });
//
type PassiveScanConfig interface {
	IPassiveScanConfig
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Parameters() IPassiveScanParameters
	SetParameters(val IPassiveScanParameters)
	Rules() *[]IPassiveScanRule
	SetRules(val *[]IPassiveScanRule)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for PassiveScanConfig
type jsiiProxy_PassiveScanConfig struct {
	jsiiProxy_IPassiveScanConfig
}

func (j *jsiiProxy_PassiveScanConfig) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanConfig) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanConfig) Parameters() IPassiveScanParameters {
	var returns IPassiveScanParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanConfig) Rules() *[]IPassiveScanRule {
	var returns *[]IPassiveScanRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanConfig) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of PassiveScanConfig.
func NewPassiveScanConfig(options IPassiveScanConfig) PassiveScanConfig {
	_init_.Initialize()

	if err := validateNewPassiveScanConfigParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PassiveScanConfig{}

	_jsii_.Create(
		"zap-cdk.PassiveScanConfig",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of PassiveScanConfig.
func NewPassiveScanConfig_Override(p PassiveScanConfig, options IPassiveScanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PassiveScanConfig",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PassiveScanConfig)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_PassiveScanConfig)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PassiveScanConfig)SetParameters(val IPassiveScanParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_PassiveScanConfig)SetRules(val *[]IPassiveScanRule) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_PassiveScanConfig)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

