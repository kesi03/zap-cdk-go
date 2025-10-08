package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a spider configuration.
type ISpider interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	Parameters() ISpiderParameters
	SetParameters(p ISpiderParameters)
	Tests() *[]ISpiderTest
	SetTests(t *[]ISpiderTest)
	Type() *string
	SetType(t *string)
}

// The jsii proxy for ISpider
type jsiiProxy_ISpider struct {
	_ byte // padding
}

func (j *jsiiProxy_ISpider) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpider)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_ISpider) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpider)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ISpider) Parameters() ISpiderParameters {
	var returns ISpiderParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpider)SetParameters(val ISpiderParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_ISpider) Tests() *[]ISpiderTest {
	var returns *[]ISpiderTest
	_jsii_.Get(
		j,
		"tests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpider)SetTests(val *[]ISpiderTest) {
	_jsii_.Set(
		j,
		"tests",
		val,
	)
}

func (j *jsiiProxy_ISpider) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISpider)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

