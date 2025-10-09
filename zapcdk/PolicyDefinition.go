package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the policy definition for an active scan.
type PolicyDefinition interface {
	IPolicyDefinition
	AlertTags() IAlertTags
	SetAlertTags(val IAlertTags)
	DefaultStrength() *string
	SetDefaultStrength(val *string)
	DefaultThreshold() *string
	SetDefaultThreshold(val *string)
	Rules() *[]IRules
	SetRules(val *[]IRules)
}

// The jsii proxy struct for PolicyDefinition
type jsiiProxy_PolicyDefinition struct {
	jsiiProxy_IPolicyDefinition
}

func (j *jsiiProxy_PolicyDefinition) AlertTags() IAlertTags {
	var returns IAlertTags
	_jsii_.Get(
		j,
		"alertTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDefinition) DefaultStrength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStrength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDefinition) DefaultThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PolicyDefinition) Rules() *[]IRules {
	var returns *[]IRules
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Creates an instance of PolicyDefinition.
func NewPolicyDefinition(options IPolicyDefinition) PolicyDefinition {
	_init_.Initialize()

	j := jsiiProxy_PolicyDefinition{}

	_jsii_.Create(
		"zap-cdk.PolicyDefinition",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of PolicyDefinition.
func NewPolicyDefinition_Override(p PolicyDefinition, options IPolicyDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PolicyDefinition",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PolicyDefinition)SetAlertTags(val IAlertTags) {
	_jsii_.Set(
		j,
		"alertTags",
		val,
	)
}

func (j *jsiiProxy_PolicyDefinition)SetDefaultStrength(val *string) {
	_jsii_.Set(
		j,
		"defaultStrength",
		val,
	)
}

func (j *jsiiProxy_PolicyDefinition)SetDefaultThreshold(val *string) {
	_jsii_.Set(
		j,
		"defaultThreshold",
		val,
	)
}

func (j *jsiiProxy_PolicyDefinition)SetRules(val *[]IRules) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

