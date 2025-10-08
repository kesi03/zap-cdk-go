package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IReplacer interface {
	DeleteAllRules() *bool
	SetDeleteAllRules(d *bool)
	Rules() *[]IReplacerRule
	SetRules(r *[]IReplacerRule)
}

// The jsii proxy for IReplacer
type jsiiProxy_IReplacer struct {
	_ byte // padding
}

func (j *jsiiProxy_IReplacer) DeleteAllRules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"deleteAllRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacer)SetDeleteAllRules(val *bool) {
	_jsii_.Set(
		j,
		"deleteAllRules",
		val,
	)
}

func (j *jsiiProxy_IReplacer) Rules() *[]IReplacerRule {
	var returns *[]IReplacerRule
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacer)SetRules(val *[]IReplacerRule) {
	if err := j.validateSetRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

