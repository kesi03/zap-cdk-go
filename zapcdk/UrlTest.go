package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing a URL test.
//
// Example:
//   const urlTest = new UrlTest({
//     name: 'test one',
//     url: 'http://www.example.com/path',
//     operator: 'and',
//     requestHeaderRegex: 'some-regex',
//     requestBodyRegex: 'some-regex',
//     responseHeaderRegex: 'some-regex',
//     responseBodyRegex: 'some-regex',
//     onFail: 'error',
//   });
//
type UrlTest interface {
	IUrlTest
	Name() *string
	SetName(val *string)
	OnFail() *string
	SetOnFail(val *string)
	Operator() *string
	SetOperator(val *string)
	RequestBodyRegex() *string
	SetRequestBodyRegex(val *string)
	RequestHeaderRegex() *string
	SetRequestHeaderRegex(val *string)
	ResponseBodyRegex() *string
	SetResponseBodyRegex(val *string)
	ResponseHeaderRegex() *string
	SetResponseHeaderRegex(val *string)
	Type() *string
	SetType(val *string)
	Url() *string
	SetUrl(val *string)
}

// The jsii proxy struct for UrlTest
type jsiiProxy_UrlTest struct {
	jsiiProxy_IUrlTest
}

func (j *jsiiProxy_UrlTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) RequestBodyRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) RequestHeaderRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestHeaderRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) ResponseBodyRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBodyRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) ResponseHeaderRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeaderRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UrlTest) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Creates an instance of UrlTest.
//
// Example:
//   const urlTest = new UrlTest({
//     name: 'test one',
//     url: 'http://www.example.com/path',
//     operator: 'and',
//     requestHeaderRegex: 'some-regex',
//     requestBodyRegex: 'some-regex',
//     responseHeaderRegex: 'some-regex',
//     responseBodyRegex: 'some-regex',
//     onFail: 'error',
//   });
//
func NewUrlTest(options IUrlTest) UrlTest {
	_init_.Initialize()

	if err := validateNewUrlTestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_UrlTest{}

	_jsii_.Create(
		"zap-cdk.UrlTest",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of UrlTest.
//
// Example:
//   const urlTest = new UrlTest({
//     name: 'test one',
//     url: 'http://www.example.com/path',
//     operator: 'and',
//     requestHeaderRegex: 'some-regex',
//     requestBodyRegex: 'some-regex',
//     responseHeaderRegex: 'some-regex',
//     responseBodyRegex: 'some-regex',
//     onFail: 'error',
//   });
//
func NewUrlTest_Override(u UrlTest, options IUrlTest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.UrlTest",
		[]interface{}{options},
		u,
	)
}

func (j *jsiiProxy_UrlTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetRequestBodyRegex(val *string) {
	_jsii_.Set(
		j,
		"requestBodyRegex",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetRequestHeaderRegex(val *string) {
	_jsii_.Set(
		j,
		"requestHeaderRegex",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetResponseBodyRegex(val *string) {
	_jsii_.Set(
		j,
		"responseBodyRegex",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetResponseHeaderRegex(val *string) {
	_jsii_.Set(
		j,
		"responseHeaderRegex",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_UrlTest)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

