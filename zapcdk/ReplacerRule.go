package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a rule for replacing strings in requests or responses.
//
// Example:
//   const replacerRule = new ReplacerRule({
//     description: 'Replace API Key',
//     url: '.*example.*',
//     matchType: 'req_header_str',
//     matchString: 'API-Key: .*',
//     matchRegex: true,
//     replacementString: 'API-Key: 12345',
//     tokenProcessing: false,
//     initiators: [1, 2, 3]
//   });
//
type ReplacerRule interface {
	IReplacerRule
	Description() *string
	SetDescription(val *string)
	Initiators() *[]*float64
	SetInitiators(val *[]*float64)
	MatchRegex() *bool
	SetMatchRegex(val *bool)
	MatchString() *string
	SetMatchString(val *string)
	MatchType() *string
	SetMatchType(val *string)
	ReplacementString() *string
	SetReplacementString(val *string)
	TokenProcessing() *bool
	SetTokenProcessing(val *bool)
	Url() *string
	SetUrl(val *string)
}

// The jsii proxy struct for ReplacerRule
type jsiiProxy_ReplacerRule struct {
	jsiiProxy_IReplacerRule
}

func (j *jsiiProxy_ReplacerRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) Initiators() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"initiators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) MatchRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"matchRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) MatchString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) ReplacementString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) TokenProcessing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"tokenProcessing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplacerRule) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Creates an instance of ReplacerRule.
//
// Example:
//   const replacerRule = new ReplacerRule({
//     description: 'Replace API Key',
//     url: '.*example.*',
//     matchType: 'req_header_str',
//     matchString: 'API-Key: .*',
//     matchRegex: true,
//     replacementString: 'API-Key: 12345',
//     tokenProcessing: false,
//     initiators: [1, 2, 3]
//   });
//
func NewReplacerRule(options IReplacerRule) ReplacerRule {
	_init_.Initialize()

	if err := validateNewReplacerRuleParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ReplacerRule{}

	_jsii_.Create(
		"zap-cdk.ReplacerRule",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ReplacerRule.
//
// Example:
//   const replacerRule = new ReplacerRule({
//     description: 'Replace API Key',
//     url: '.*example.*',
//     matchType: 'req_header_str',
//     matchString: 'API-Key: .*',
//     matchRegex: true,
//     replacementString: 'API-Key: 12345',
//     tokenProcessing: false,
//     initiators: [1, 2, 3]
//   });
//
func NewReplacerRule_Override(r ReplacerRule, options IReplacerRule) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ReplacerRule",
		[]interface{}{options},
		r,
	)
}

func (j *jsiiProxy_ReplacerRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetInitiators(val *[]*float64) {
	_jsii_.Set(
		j,
		"initiators",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetMatchRegex(val *bool) {
	_jsii_.Set(
		j,
		"matchRegex",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetMatchString(val *string) {
	if err := j.validateSetMatchStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchString",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetMatchType(val *string) {
	if err := j.validateSetMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchType",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetReplacementString(val *string) {
	if err := j.validateSetReplacementStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementString",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetTokenProcessing(val *bool) {
	_jsii_.Set(
		j,
		"tokenProcessing",
		val,
	)
}

func (j *jsiiProxy_ReplacerRule)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

