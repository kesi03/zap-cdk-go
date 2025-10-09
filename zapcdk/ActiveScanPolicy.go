package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing an active scan policy configuration.
type ActiveScanPolicy interface {
	IActiveScanPolicy
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Parameters() IActiveScanPolicyParameters
	SetParameters(val IActiveScanPolicyParameters)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for ActiveScanPolicy
type jsiiProxy_ActiveScanPolicy struct {
	jsiiProxy_IActiveScanPolicy
}

func (j *jsiiProxy_ActiveScanPolicy) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicy) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicy) Parameters() IActiveScanPolicyParameters {
	var returns IActiveScanPolicyParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveScanPolicy) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of ActiveScanPolicy.
func NewActiveScanPolicy(options IActiveScanPolicy) ActiveScanPolicy {
	_init_.Initialize()

	if err := validateNewActiveScanPolicyParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveScanPolicy{}

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicy",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ActiveScanPolicy.
func NewActiveScanPolicy_Override(a ActiveScanPolicy, options IActiveScanPolicy) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ActiveScanPolicy",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_ActiveScanPolicy)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicy)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicy)SetParameters(val IActiveScanPolicyParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ActiveScanPolicy)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

