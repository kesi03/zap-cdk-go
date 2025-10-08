package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPolicyDefinition interface {
	AlertTags() IAlertTags
	SetAlertTags(a IAlertTags)
	DefaultStrength() *string
	SetDefaultStrength(d *string)
	DefaultThreshold() *string
	SetDefaultThreshold(d *string)
	Rules() *[]IRules
	SetRules(r *[]IRules)
}

// The jsii proxy for IPolicyDefinition
type jsiiProxy_IPolicyDefinition struct {
	_ byte // padding
}

func (j *jsiiProxy_IPolicyDefinition) AlertTags() IAlertTags {
	var returns IAlertTags
	_jsii_.Get(
		j,
		"alertTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyDefinition)SetAlertTags(val IAlertTags) {
	_jsii_.Set(
		j,
		"alertTags",
		val,
	)
}

func (j *jsiiProxy_IPolicyDefinition) DefaultStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyDefinition)SetDefaultStrength(val *string) {
	_jsii_.Set(
		j,
		"defaultStrength",
		val,
	)
}

func (j *jsiiProxy_IPolicyDefinition) DefaultThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyDefinition)SetDefaultThreshold(val *string) {
	_jsii_.Set(
		j,
		"defaultThreshold",
		val,
	)
}

func (j *jsiiProxy_IPolicyDefinition) Rules() *[]IRules {
	var returns *[]IRules
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPolicyDefinition)SetRules(val *[]IRules) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

