package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IOpenAPI interface {
	ApiFile() *string
	SetApiFile(a *string)
	ApiUrl() *string
	SetApiUrl(a *string)
	Context() *string
	SetContext(c *string)
	TargetUrl() *string
	SetTargetUrl(t *string)
	User() *string
	SetUser(u *string)
}

// The jsii proxy for IOpenAPI
type jsiiProxy_IOpenAPI struct {
	_ byte // padding
}

func (j *jsiiProxy_IOpenAPI) ApiFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenAPI)SetApiFile(val *string) {
	_jsii_.Set(
		j,
		"apiFile",
		val,
	)
}

func (j *jsiiProxy_IOpenAPI) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenAPI)SetApiUrl(val *string) {
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

func (j *jsiiProxy_IOpenAPI) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenAPI)SetContext(val *string) {
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_IOpenAPI) TargetUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenAPI)SetTargetUrl(val *string) {
	_jsii_.Set(
		j,
		"targetUrl",
		val,
	)
}

func (j *jsiiProxy_IOpenAPI) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IOpenAPI)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

