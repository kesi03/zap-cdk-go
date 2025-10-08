package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IActiveScanConfig interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	ExcludePaths() *[]*string
	SetExcludePaths(e *[]*string)
	Parameters() IActiveScanConfigParameters
	SetParameters(p IActiveScanConfigParameters)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IActiveScanConfig
type jsiiProxy_IActiveScanConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IActiveScanConfig) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfig)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfig) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfig)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfig) ExcludePaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfig)SetExcludePaths(val *[]*string) {
	_jsii_.Set(
		j,
		"excludePaths",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfig) Parameters() IActiveScanConfigParameters {
	var returns IActiveScanConfigParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfig)SetParameters(val IActiveScanConfigParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IActiveScanConfig) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IActiveScanConfig)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

