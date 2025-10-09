package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the parameters for authentication in the scanning process.
type AuthenticationParameters interface {
	IAuthenticationParameters
	Method() *string
	SetMethod(val *string)
	Parameters() IAuthenticationParametersParameters
	SetParameters(val IAuthenticationParametersParameters)
	Verification() IAuthenticationParametersVerification
	SetVerification(val IAuthenticationParametersVerification)
}

// The jsii proxy struct for AuthenticationParameters
type jsiiProxy_AuthenticationParameters struct {
	jsiiProxy_IAuthenticationParameters
}

func (j *jsiiProxy_AuthenticationParameters) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParameters) Parameters() IAuthenticationParametersParameters {
	var returns IAuthenticationParametersParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuthenticationParameters) Verification() IAuthenticationParametersVerification {
	var returns IAuthenticationParametersVerification
	_jsii_.Get(
		j,
		"verification",
		&returns,
	)
	return returns
}


// Creates an instance of AuthenticationParameters.
func NewAuthenticationParameters(options IAuthenticationParameters) AuthenticationParameters {
	_init_.Initialize()

	if err := validateNewAuthenticationParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuthenticationParameters{}

	_jsii_.Create(
		"zap-cdk.AuthenticationParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AuthenticationParameters.
func NewAuthenticationParameters_Override(a AuthenticationParameters, options IAuthenticationParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AuthenticationParameters",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AuthenticationParameters)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParameters)SetParameters(val IAuthenticationParametersParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_AuthenticationParameters)SetVerification(val IAuthenticationParametersVerification) {
	if err := j.validateSetVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verification",
		val,
	)
}

