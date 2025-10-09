package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

type RequestorParameters interface {
	IRequestorParameters
	AlwaysRun() *bool
	SetAlwaysRun(val *bool)
	Enabled() *bool
	SetEnabled(val *bool)
	Requests() *[]IRequest
	SetRequests(val *[]IRequest)
	User() *string
	SetUser(val *string)
}

// The jsii proxy struct for RequestorParameters
type jsiiProxy_RequestorParameters struct {
	jsiiProxy_IRequestorParameters
}

func (j *jsiiProxy_RequestorParameters) AlwaysRun() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"alwaysRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestorParameters) Enabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestorParameters) Requests() *[]IRequest {
	var returns *[]IRequest
	_jsii_.Get(
		j,
		"requests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RequestorParameters) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


// Creates an instance of RequestorParameters.
//
// Example:
//   const requestorParams = new RequestorParameters({
//     user: 'admin',
//     requests: [
//       new Request({ url: 'https://example.com/api1' }),
//       new Request({ url: 'https://example.com/api2', method: 'POST' })
//     ],
//     enabled: true,
//     alwaysRun: false
//   });
//
func NewRequestorParameters(options IRequestorParameters) RequestorParameters {
	_init_.Initialize()

	if err := validateNewRequestorParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_RequestorParameters{}

	_jsii_.Create(
		"zap-cdk.RequestorParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of RequestorParameters.
//
// Example:
//   const requestorParams = new RequestorParameters({
//     user: 'admin',
//     requests: [
//       new Request({ url: 'https://example.com/api1' }),
//       new Request({ url: 'https://example.com/api2', method: 'POST' })
//     ],
//     enabled: true,
//     alwaysRun: false
//   });
//
func NewRequestorParameters_Override(r RequestorParameters, options IRequestorParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.RequestorParameters",
		[]interface{}{options},
		r,
	)
}

func (j *jsiiProxy_RequestorParameters)SetAlwaysRun(val *bool) {
	_jsii_.Set(
		j,
		"alwaysRun",
		val,
	)
}

func (j *jsiiProxy_RequestorParameters)SetEnabled(val *bool) {
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_RequestorParameters)SetRequests(val *[]IRequest) {
	if err := j.validateSetRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requests",
		val,
	)
}

func (j *jsiiProxy_RequestorParameters)SetUser(val *string) {
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

