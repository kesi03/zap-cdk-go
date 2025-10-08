package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a single request configuration.
type IRequest interface {
	Data() *string
	SetData(d *string)
	Headers() *[]*string
	SetHeaders(h *[]*string)
	HttpVersion() *string
	SetHttpVersion(h *string)
	Method() *string
	SetMethod(m *string)
	Name() *string
	SetName(n *string)
	ResponseCode() *float64
	SetResponseCode(r *float64)
	Url() *string
	SetUrl(u *string)
}

// The jsii proxy for IRequest
type jsiiProxy_IRequest struct {
	_ byte // padding
}

func (j *jsiiProxy_IRequest) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetData(val *string) {
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_IRequest) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetHeaders(val *[]*string) {
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_IRequest) HttpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetHttpVersion(val *string) {
	_jsii_.Set(
		j,
		"httpVersion",
		val,
	)
}

func (j *jsiiProxy_IRequest) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetMethod(val *string) {
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_IRequest) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IRequest) ResponseCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"responseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetResponseCode(val *float64) {
	_jsii_.Set(
		j,
		"responseCode",
		val,
	)
}

func (j *jsiiProxy_IRequest) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequest)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

