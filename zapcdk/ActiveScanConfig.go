package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for an active scan.
//
// Example:
//   const activeScanConfig = new ActiveScanConfig({
//     parameters: new ActiveScanConfigParameters(),
//     excludePaths: ['^/api/health$', '^/static/.*$'],
//     enabled: true,
//     alwaysRun: false
//   });
//   console.log(activeScanConfig.type); // 'activeScan-config'
//
type ActiveScanConfig interface {
	IActiveScanConfig
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	ExcludePaths() *[]*string
	SetExcludePaths(val *[]*string)
	Parameters() IActiveScanConfigParameters
	SetParameters(val IActiveScanConfigParameters)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for ActiveScanConfig
type jsiiProxy_ActiveScanConfig struct {
	jsiiProxy_IActiveScanConfig
}

func (j *jsiiProxy_ActiveScanConfig) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfig) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfig) ExcludePaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfig) Parameters() IActiveScanConfigParameters {
	var returns IActiveScanConfigParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanConfig) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanConfig.
//
// Example:
//   const activeScanConfig = new ActiveScanConfig({
//     parameters: new ActiveScanConfigParameters(),
//     excludePaths: ['^/api/health$', '^/static/.*$'],
//     enabled: true,
//     alwaysRun: false
//   });
//   console.log(activeScanConfig.type); // 'activeScan-config'
//
func NewActiveScanConfig(options IActiveScanConfig) ActiveScanConfig {
	_init_.Initialize()

	if err := validateNewActiveScanConfigParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanConfig{}

	_jsii_.Create(
		"zap-cdk.ActiveScanConfig",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanConfig.
//
// Example:
//   const activeScanConfig = new ActiveScanConfig({
//     parameters: new ActiveScanConfigParameters(),
//     excludePaths: ['^/api/health$', '^/static/.*$'],
//     enabled: true,
//     alwaysRun: false
//   });
//   console.log(activeScanConfig.type); // 'activeScan-config'
//
func NewActiveScanConfig_Override(a ActiveScanConfig, options IActiveScanConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanConfig",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScanConfig)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfig)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfig)SetExcludePaths(val *[]*string) {
	_jsii_.Set(
		j,
		"excludePaths",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfig)SetParameters(val IActiveScanConfigParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ActiveScanConfig)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

