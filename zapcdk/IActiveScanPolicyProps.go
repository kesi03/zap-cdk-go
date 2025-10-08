package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ActiveScanPolicyConfig construct.
type IActiveScanPolicyProps interface {
	ActiveScanPolicy() IActiveScanPolicy
	SetActiveScanPolicy(a IActiveScanPolicy)
}

// The jsii proxy for IActiveScanPolicyProps
type jsiiProxy_IActiveScanPolicyProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanPolicyProps) ActiveScanPolicy() IActiveScanPolicy {
	var returns IActiveScanPolicy
	_jsii_.Get(
		j,
		"activeScanPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanPolicyProps)SetActiveScanPolicy(val IActiveScanPolicy) {
	if err := j.validateSetActiveScanPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeScanPolicy",
		val,
	)
}

