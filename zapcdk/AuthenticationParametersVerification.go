package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the verification details for authentication in the scanning process.
type AuthenticationParametersVerification interface {
	IAuthenticationParametersVerification
	LoggedInRegex() *string
	SetLoggedInRegex(val *string)
	LoggedOutRegex() *string
	SetLoggedOutRegex(val *string)
	Method() *string
	SetMethod(val *string)
	PollAdditionalHeaders() *[]IPollAdditionalHeaders
	SetPollAdditionalHeaders(val *[]IPollAdditionalHeaders)
	PollFrequency() *float64
	SetPollFrequency(val *float64)
	PollPostData() *string
	SetPollPostData(val *string)
	PollUnits() *string
	SetPollUnits(val *string)
	PollUrl() *string
	SetPollUrl(val *string)
}

// The jsii proxy struct for AuthenticationParametersVerification
type jsiiProxy_AuthenticationParametersVerification struct {
	jsiiProxy_IAuthenticationParametersVerification
}

func (j *jsiiProxy_AuthenticationParametersVerification) LoggedInRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggedInRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) LoggedOutRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggedOutRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) PollAdditionalHeaders() *[]IPollAdditionalHeaders {
	var returns *[]IPollAdditionalHeaders
	_jsii_.Get(
		j,
		"pollAdditionalHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) PollFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pollFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) PollPostData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollPostData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) PollUnits() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParametersVerification) PollUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pollUrl",
		&returns,
	)
	return returns
}


// Creates an instance of AuthenticationParametersVerification.
func NewAuthenticationParametersVerification(options IAuthenticationParametersVerification) AuthenticationParametersVerification {
	_init_.Initialize()

	if err := validateNewAuthenticationParametersVerificationParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthenticationParametersVerification{}

	_jsii_.Create(
		"zap-cdk.AuthenticationParametersVerification",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AuthenticationParametersVerification.
func NewAuthenticationParametersVerification_Override(a AuthenticationParametersVerification, options IAuthenticationParametersVerification) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AuthenticationParametersVerification",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetLoggedInRegex(val *string) {
	_jsii_.Set(
		j,
		"loggedInRegex",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetLoggedOutRegex(val *string) {
	_jsii_.Set(
		j,
		"loggedOutRegex",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetPollAdditionalHeaders(val *[]IPollAdditionalHeaders) {
	_jsii_.Set(
		j,
		"pollAdditionalHeaders",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetPollFrequency(val *float64) {
	_jsii_.Set(
		j,
		"pollFrequency",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetPollPostData(val *string) {
	_jsii_.Set(
		j,
		"pollPostData",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetPollUnits(val *string) {
	_jsii_.Set(
		j,
		"pollUnits",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParametersVerification)SetPollUrl(val *string) {
	_jsii_.Set(
		j,
		"pollUrl",
		val,
	)
}

