package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing the parameters for making requests.
type IRequestorParameters interface {
	AlwaysRun() *bool
	SetAlwaysRun(a *bool)
	Enabled() *bool
	SetEnabled(e *bool)
	Requests() *[]IRequest
	SetRequests(r *[]IRequest)
	User() *string
	SetUser(u *string)
}

// The jsii proxy for IRequestorParameters
type jsiiProxy_IRequestorParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IRequestorParameters) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestorParameters)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_IRequestorParameters) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestorParameters)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_IRequestorParameters) Requests() *[]IRequest {
	var returns *[]IRequest
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestorParameters)SetRequests(val *[]IRequest) {
	if err := j.validateSetRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requests",
		val,
	)
}

func (j *jsiiProxy_IRequestorParameters) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRequestorParameters)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

