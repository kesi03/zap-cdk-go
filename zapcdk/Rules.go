package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a rule for the active scan.
type Rules interface {
	IRules
	Id() *float64
	SetId(val *float64)
	Name() *string
	SetName(val *string)
	Strength() *string
	SetStrength(val *string)
	Threshold() *string
	SetThreshold(val *string)
}

// The jsii proxy struct for Rules
type jsiiProxy_Rules struct {
	jsiiProxy_IRules
}

func (j *jsiiProxy_Rules) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rules) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rules) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Rules) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}


// Creates an instance of Rules.
func NewRules(options IRules) Rules {
	_init_.Initialize()

	if err := validateNewRulesParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Rules{}

	_jsii_.Create(
		"zap-cdk.Rules",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Rules.
func NewRules_Override(r Rules, options IRules) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Rules",
		[]interface{}{options},
		r,
	)
}

func (j *jsiiProxy_Rules)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Rules)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Rules)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_Rules)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

