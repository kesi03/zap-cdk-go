package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPassiveScanConfig interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	Parameters() IPassiveScanParameters
	SetParameters(p IPassiveScanParameters)
	Rules() *[]IPassiveScanRule
	SetRules(r *[]IPassiveScanRule)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for IPassiveScanConfig
type jsiiProxy_IPassiveScanConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IPassiveScanConfig) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanConfig)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanConfig) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanConfig)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanConfig) Parameters() IPassiveScanParameters {
	var returns IPassiveScanParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanConfig)SetParameters(val IPassiveScanParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanConfig) Rules() *[]IPassiveScanRule {
	var returns *[]IPassiveScanRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanConfig)SetRules(val *[]IPassiveScanRule) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_IPassiveScanConfig) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPassiveScanConfig)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

