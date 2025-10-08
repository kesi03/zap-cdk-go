package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScan interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	Parameters() IActiveScanParameters
	SetParameters(p IActiveScanParameters)
	PolicyDefinition() IPolicyDefinition
	SetPolicyDefinition(p IPolicyDefinition)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IActiveScan
type jsiiProxy_IActiveScan struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScan) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScan)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_IActiveScan) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScan)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IActiveScan) Parameters() IActiveScanParameters {
	var returns IActiveScanParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScan)SetParameters(val IActiveScanParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IActiveScan) PolicyDefinition() IPolicyDefinition {
	var returns IPolicyDefinition
	_jsii_.Get(
		j,
		"policyDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScan)SetPolicyDefinition(val IPolicyDefinition) {
	_jsii_.Set(
		j,
		"policyDefinition",
		val,
	)
}

func (j *jsiiProxy_IActiveScan) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScan)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

