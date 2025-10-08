package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Properties for the PassiveScanConfig construct.
type IPassiveScanConfigProps interface {
	PassiveScanConfig() IPassiveScanConfig
	SetPassiveScanConfig(p IPassiveScanConfig)
}

// The jsii proxy for IPassiveScanConfigProps
type jsiiProxy_IPassiveScanConfigProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IPassiveScanConfigProps) PassiveScanConfig() IPassiveScanConfig {
	var returns IPassiveScanConfig
	_jsii_.Get(
		j,
		"passiveScanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanConfigProps)SetPassiveScanConfig(val IPassiveScanConfig) {
	if err := j.validateSetPassiveScanConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveScanConfig",
		val,
	)
}

