package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IReplacerRule interface {
	Description() *string
	SetDescription(d *string)
	Initiators() *[]*float64
	SetInitiators(i *[]*float64)
	MatchRegex() *bool
	SetMatchRegex(m *bool)
	MatchString() *string
	SetMatchString(m *string)
	MatchType() *string
	SetMatchType(m *string)
	ReplacementString() *string
	SetReplacementString(r *string)
	TokenProcessing() *bool
	SetTokenProcessing(t *bool)
	Url() *string
	SetUrl(u *string)
}

// The jsii proxy for IReplacerRule
type jsiiProxy_IReplacerRule struct {
	_ byte // padding
}

func (j *jsiiProxy_IReplacerRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) Initiators() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"initiators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetInitiators(val *[]*float64) {
	_jsii_.Set(
		j,
		"initiators",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) MatchRegex() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"matchRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetMatchRegex(val *bool) {
	_jsii_.Set(
		j,
		"matchRegex",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) MatchString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetMatchString(val *string) {
	if err := j.validateSetMatchStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchString",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) MatchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetMatchType(val *string) {
	if err := j.validateSetMatchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchType",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) ReplacementString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replacementString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetReplacementString(val *string) {
	if err := j.validateSetReplacementStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementString",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) TokenProcessing() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"tokenProcessing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetTokenProcessing(val *bool) {
	_jsii_.Set(
		j,
		"tokenProcessing",
		val,
	)
}

func (j *jsiiProxy_IReplacerRule) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplacerRule)SetUrl(val *string) {
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

