package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a spider configuration.
//
// Example:
//   const spiderParams = new SpiderParameters({
//     context: 'MyContext',
//     maxDuration: 10,
//     parseComments: false
//   });
//   const spiderTest = new SpiderTest({
//     statistic: 'urls',
//     operator: '>=',
//     value: 10,
//     onFail: 'error'
//   });
//   const spider = new Spider(spiderParams, [spiderTest], true, false);
//
type Spider interface {
	ISpider
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Parameters() ISpiderParameters
	SetParameters(val ISpiderParameters)
	Tests() *[]ISpiderTest
	SetTests(val *[]ISpiderTest)
	Type() *string
	SetType(val *string)
}

// The jsii proxy struct for Spider
type jsiiProxy_Spider struct {
	jsiiProxy_ISpider
}

func (j *jsiiProxy_Spider) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Spider) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Spider) Parameters() ISpiderParameters {
	var returns ISpiderParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Spider) Tests() *[]ISpiderTest {
	var returns *[]ISpiderTest
	_jsii_.Get(
		j,
		"tests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Spider) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Creates an instance of Spider.
func NewSpider(parameters ISpiderParameters, tests *[]ISpiderTest, enabled *bool, alwaysRun *bool) Spider {
	_init_.Initialize()

	if err := validateNewSpiderParameters(parameters); err != nil {
		panic(err)
	}
	j := jsiiProxy_Spider{}

	_jsii_.Create(
		"zap-cdk.Spider",
		[]interface{}{parameters, tests, enabled, alwaysRun},
		&j,
	)

	return &j
}

// Creates an instance of Spider.
func NewSpider_Override(s Spider, parameters ISpiderParameters, tests *[]ISpiderTest, enabled *bool, alwaysRun *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Spider",
		[]interface{}{parameters, tests, enabled, alwaysRun},
		s,
	)
}

func (j *jsiiProxy_Spider)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_Spider)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_Spider)SetParameters(val ISpiderParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_Spider)SetTests(val *[]ISpiderTest) {
	_jsii_.Set(
		j,
		"tests",
		val,
	)
}

func (j *jsiiProxy_Spider)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

