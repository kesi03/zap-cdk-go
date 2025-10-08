package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for URL tests.
//
// Example YAML representation:
// ```yaml
// - name: 'test one'                      # Name of the test, optional
//   type: url                             # Specifies that the test is of type 'url'
//   url: http://www.example.com/path      # String: The URL to be tested.
//   operator: 'and'                       # One of 'and', 'or', default is 'or'
//   requestHeaderRegex:                   # String: The regular expression to be matched in the request header, optional
//   requestBodyRegex:                     # String: The regular expression to be matched in the request body, optional
//   responseHeaderRegex:                  # String: The regular expression to be matched in the response header, optional
//   responseBodyRegex:                    # String: The regular expression to be matched in the response body, optional
//   onFail: 'info'                        # String: One of 'warn', 'error', 'info', mandatory
// ```.
type IUrlTest interface {
	Name() *string
	SetName(n *string)
	OnFail() *string
	SetOnFail(o *string)
	Operator() *string
	SetOperator(o *string)
	RequestBodyRegex() *string
	SetRequestBodyRegex(r *string)
	RequestHeaderRegex() *string
	SetRequestHeaderRegex(r *string)
	ResponseBodyRegex() *string
	SetResponseBodyRegex(r *string)
	ResponseHeaderRegex() *string
	SetResponseHeaderRegex(r *string)
	Type() *string
	SetType(t *string)
	Url() *string
	SetUrl(u *string)
}

// The jsii proxy for IUrlTest
type jsiiProxy_IUrlTest struct {
	_ byte // padding
}

func (j *jsiiProxy_IUrlTest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) OnFail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onFail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetOnFail(val *string) {
	if err := j.validateSetOnFailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onFail",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetOperator(val *string) {
	if err := j.validateSetOperatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operator",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) RequestBodyRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetRequestBodyRegex(val *string) {
	_jsii_.Set(
		j,
		"requestBodyRegex",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) RequestHeaderRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestHeaderRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetRequestHeaderRegex(val *string) {
	_jsii_.Set(
		j,
		"requestHeaderRegex",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) ResponseBodyRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseBodyRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetResponseBodyRegex(val *string) {
	_jsii_.Set(
		j,
		"responseBodyRegex",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) ResponseHeaderRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseHeaderRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetResponseHeaderRegex(val *string) {
	_jsii_.Set(
		j,
		"responseHeaderRegex",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IUrlTest) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUrlTest)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

