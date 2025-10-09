package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type Request interface {
	IRequest
	Data() *string
	SetData(val *string)
	Headers() *[]*string
	SetHeaders(val *[]*string)
	HttpVersion() *string
	SetHttpVersion(val *string)
	Method() *string
	SetMethod(val *string)
	Name() *string
	SetName(val *string)
	ResponseCode() *float64
	SetResponseCode(val *float64)
	Url() *string
	SetUrl(val *string)
}

// The jsii proxy struct for Request
type jsiiProxy_Request struct {
	jsiiProxy_IRequest
}

func (j *jsiiProxy_Request) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Request) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Request) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Request) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Request) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Request) ResponseCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Request) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


// Creates an instance of Request.
//
// Example:
//   const request = new Request({
//     url: 'https://example.com/api',
//     method: 'POST',
//     headers: ['Content-Type: application/json'],
//     data: '{"key":"value"}',
//     responseCode: 200
//   });
//
func NewRequest(options IRequest) Request {
	_init_.Initialize()

	if err := validateNewRequestParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Request{}

	_jsii_.Create(
		"zap-cdk.Request",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Request.
//
// Example:
//   const request = new Request({
//     url: 'https://example.com/api',
//     method: 'POST',
//     headers: ['Content-Type: application/json'],
//     data: '{"key":"value"}',
//     responseCode: 200
//   });
//
func NewRequest_Override(r Request, options IRequest) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Request",
		[]interface{}{options},
		r,
	)
}

func (j *jsiiProxy_Request)SetData(val *string) {
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_Request)SetHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_Request)SetHttpVersion(val *string) {
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_Request)SetMethod(val *string) {
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_Request)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Request)SetResponseCode(val *float64) {
	_jsii_.Set(
		j,
		"responseCode",
		val,
	)
}

func (j *jsiiProxy_Request)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

