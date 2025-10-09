package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing an active scan configuration.
type ActiveScan interface {
	IActiveScan
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Parameters() IActiveScanParameters
	SetParameters(val IActiveScanParameters)
	PolicyDefinition() IPolicyDefinition
	SetPolicyDefinition(val IPolicyDefinition)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for ActiveScan
type jsiiProxy_ActiveScan struct {
	jsiiProxy_IActiveScan
}

func (j *jsiiProxy_ActiveScan) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScan) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScan) Parameters() IActiveScanParameters {
	var returns IActiveScanParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScan) PolicyDefinition() IPolicyDefinition {
	var returns IPolicyDefinition
	_jsii_.Get(
		j,
		"policyDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScan) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScan.
func NewActiveScan(options IActiveScan) ActiveScan {
	_init_.Initialize()

	if err := validateNewActiveScanParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScan{}

	_jsii_.Create(
		"zap-cdk.ActiveScan",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScan.
func NewActiveScan_Override(a ActiveScan, options IActiveScan) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScan",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScan)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_ActiveScan)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveScan)SetParameters(val IActiveScanParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ActiveScan)SetPolicyDefinition(val IPolicyDefinition) {
	_jsii_.Set(
		j,
		"policyDefinition",
		val,
	)
}

func (j *jsiiProxy_ActiveScan)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

