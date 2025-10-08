package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAuthenticationParametersVerification interface {
	LoggedInRegex() *string
	SetLoggedInRegex(l *string)
	LoggedOutRegex() *string
	SetLoggedOutRegex(l *string)
	Method() *string
	SetMethod(m *string)
	PollAdditionalHeaders() *[]IPollAdditionalHeaders
	SetPollAdditionalHeaders(p *[]IPollAdditionalHeaders)
	PollFrequency() *float64
	SetPollFrequency(p *float64)
	PollPostData() *string
	SetPollPostData(p *string)
	PollUnits() *string
	SetPollUnits(p *string)
	PollUrl() *string
	SetPollUrl(p *string)
}

// The jsii proxy for IAuthenticationParametersVerification
type jsiiProxy_IAuthenticationParametersVerification struct {
	_ byte // padding
}

func (j *jsiiProxy_IAuthenticationParametersVerification) LoggedInRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggedInRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetLoggedInRegex(val *string) {
	_jsii_.Set(
		j,
		"loggedInRegex",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) LoggedOutRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggedOutRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetLoggedOutRegex(val *string) {
	_jsii_.Set(
		j,
		"loggedOutRegex",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) PollAdditionalHeaders() *[]IPollAdditionalHeaders {
	var returns *[]IPollAdditionalHeaders
	_jsii_.Get(
		j,
		"pollAdditionalHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetPollAdditionalHeaders(val *[]IPollAdditionalHeaders) {
	_jsii_.Set(
		j,
		"pollAdditionalHeaders",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) PollFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pollFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetPollFrequency(val *float64) {
	_jsii_.Set(
		j,
		"pollFrequency",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) PollPostData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollPostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetPollPostData(val *string) {
	_jsii_.Set(
		j,
		"pollPostData",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) PollUnits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetPollUnits(val *string) {
	_jsii_.Set(
		j,
		"pollUnits",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParametersVerification) PollUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParametersVerification)SetPollUrl(val *string) {
	_jsii_.Set(
		j,
		"pollUrl",
		val,
	)
}

