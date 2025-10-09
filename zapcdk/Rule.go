package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing an individual rule in the active scan policy.
type Rule interface {
	IRule
	Id() *float64
	SetId(val *float64)
	Name() *string
	SetName(val *string)
	Strength() *string
	SetStrength(val *string)
	Threshold() *string
	SetThreshold(val *string)
}

// The jsii proxy struct for Rule
type jsiiProxy_Rule struct {
	jsiiProxy_IRule
}

func (j *jsiiProxy_Rule) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rule) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}


// Creates an instance of Rule.
func NewRule(options IRule) Rule {
	_init_.Initialize()

	if err := validateNewRuleParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Rule{}

	_jsii_.Create(
		"zap-cdk.Rule",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Rule.
func NewRule_Override(r Rule, options IRule) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Rule",
		[]interface{}{options},
		r,
	)
}

func (j *jsiiProxy_Rule)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Rule)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Rule)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_Rule)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

