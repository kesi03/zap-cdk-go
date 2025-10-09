package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for string replacement rules.
//
// Example:
//   const replacerConfig = new Replacer({
//     deleteAllRules: true,
//     rules: [
//       new ReplacerRule({
//         description: 'Replace API Key',
//         url: '.*example.*',
//         matchType: 'req_header_str',
//         matchString: 'API-Key: .*',
//         matchRegex: true,
//         replacementString: 'API-Key: 12345',
//         tokenProcessing: false,
//         initiators: [1, 2, 3]
//       })
//     ]
//   });
//
type Replacer interface {
	IReplacer
	DeleteAllRules() *bool
	SetDeleteAllRules(val *bool)
	Rules() *[]IReplacerRule
	SetRules(val *[]IReplacerRule)
}

// The jsii proxy struct for Replacer
type jsiiProxy_Replacer struct {
	jsiiProxy_IReplacer
}

func (j *jsiiProxy_Replacer) DeleteAllRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deleteAllRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Replacer) Rules() *[]IReplacerRule {
	var returns *[]IReplacerRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}


// Creates an instance of Replacer.
func NewReplacer(options IReplacer) Replacer {
	_init_.Initialize()

	if err := validateNewReplacerParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Replacer{}

	_jsii_.Create(
		"zap-cdk.Replacer",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Replacer.
func NewReplacer_Override(r Replacer, options IReplacer) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Replacer",
		[]interface{}{options},
		r,
	)
}

func (j *jsiiProxy_Replacer)SetDeleteAllRules(val *bool) {
	_jsii_.Set(
		j,
		"deleteAllRules",
		val,
	)
}

func (j *jsiiProxy_Replacer)SetRules(val *[]IReplacerRule) {
	if err := j.validateSetRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

