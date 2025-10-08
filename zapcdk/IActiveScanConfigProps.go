package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the ActiveScanConfig construct.
type IActiveScanConfigProps interface {
	ActiveScanConfig() IActiveScanConfig
	SetActiveScanConfig(a IActiveScanConfig)
}

// The jsii proxy for IActiveScanConfigProps
type jsiiProxy_IActiveScanConfigProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanConfigProps) ActiveScanConfig() IActiveScanConfig {
	var returns IActiveScanConfig
	_jsii_.Get(
		j,
		"activeScanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfigProps)SetActiveScanConfig(val IActiveScanConfig) {
	if err := j.validateSetActiveScanConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeScanConfig",
		val,
	)
}

