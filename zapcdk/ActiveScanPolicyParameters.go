package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the parameters for an active scan policy.
type ActiveScanPolicyParameters interface {
	IActiveScanPolicyParameters
	Name() *string
	SetName(val *string)
	PolicyDefinition() IActiveScanPolicyDefinition
	SetPolicyDefinition(val IActiveScanPolicyDefinition)
}

// The jsii proxy struct for ActiveScanPolicyParameters
type jsiiProxy_ActiveScanPolicyParameters struct {
	jsiiProxy_IActiveScanPolicyParameters
}

func (j *jsiiProxy_ActiveScanPolicyParameters) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicyParameters) PolicyDefinition() IActiveScanPolicyDefinition {
	var returns IActiveScanPolicyDefinition
	_jsii_.Get(
		j,
		"policyDefinition",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanPolicyParameters.
func NewActiveScanPolicyParameters(options IActiveScanPolicyParameters) ActiveScanPolicyParameters {
	_init_.Initialize()

	if err := validateNewActiveScanPolicyParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanPolicyParameters{}

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicyParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanPolicyParameters.
func NewActiveScanPolicyParameters_Override(a ActiveScanPolicyParameters, options IActiveScanPolicyParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicyParameters",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScanPolicyParameters)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicyParameters)SetPolicyDefinition(val IActiveScanPolicyDefinition) {
	if err := j.validateSetPolicyDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDefinition",
		val,
	)
}

