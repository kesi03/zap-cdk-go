package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScanPolicyParameters interface {
	Name() *string
	SetName(n *string)
	PolicyDefinition() IActiveScanPolicyDefinition
	SetPolicyDefinition(p IActiveScanPolicyDefinition)
}

// The jsii proxy for IActiveScanPolicyParameters
type jsiiProxy_IActiveScanPolicyParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanPolicyParameters) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyParameters)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IActiveScanPolicyParameters) PolicyDefinition() IActiveScanPolicyDefinition {
	var returns IActiveScanPolicyDefinition
	_jsii_.Get(
		j,
		"policyDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyParameters)SetPolicyDefinition(val IActiveScanPolicyDefinition) {
	if err := j.validateSetPolicyDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDefinition",
		val,
	)
}

