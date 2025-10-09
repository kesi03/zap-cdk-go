package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a passive scan rule configuration.
//
// Example:
//   const passiveScanRule = new PassiveScanRule({
//     id: 10010,
//     name: 'Cross-Domain Misconfiguration',
//     threshold: 'Medium'
//   });
//
type PassiveScanRule interface {
	IPassiveScanRule
	Id() *float64
	SetId(val *float64)
	Name() *string
	SetName(val *string)
	Threshold() *string
	SetThreshold(val *string)
}

// The jsii proxy struct for PassiveScanRule
type jsiiProxy_PassiveScanRule struct {
	jsiiProxy_IPassiveScanRule
}

func (j *jsiiProxy_PassiveScanRule) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PassiveScanRule) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}


// Creates an instance of PassiveScanRule.
func NewPassiveScanRule(options IPassiveScanRule) PassiveScanRule {
	_init_.Initialize()

	if err := validateNewPassiveScanRuleParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_PassiveScanRule{}

	_jsii_.Create(
		"zap-cdk.PassiveScanRule",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of PassiveScanRule.
func NewPassiveScanRule_Override(p PassiveScanRule, options IPassiveScanRule) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.PassiveScanRule",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_PassiveScanRule)SetId(val *float64) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PassiveScanRule)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PassiveScanRule)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

