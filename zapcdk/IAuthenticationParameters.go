package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAuthenticationParameters interface {
	Method() *string
	SetMethod(m *string)
	Parameters() IAuthenticationParametersParameters
	SetParameters(p IAuthenticationParametersParameters)
	Verification() IAuthenticationParametersVerification
	SetVerification(v IAuthenticationParametersVerification)
}

// The jsii proxy for IAuthenticationParameters
type jsiiProxy_IAuthenticationParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_IAuthenticationParameters) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParameters)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParameters) Parameters() IAuthenticationParametersParameters {
	var returns IAuthenticationParametersParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParameters)SetParameters(val IAuthenticationParametersParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IAuthenticationParameters) Verification() IAuthenticationParametersVerification {
	var returns IAuthenticationParametersVerification
	_jsii_.Get(
		j,
		"verification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAuthenticationParameters)SetVerification(val IAuthenticationParametersVerification) {
	if err := j.validateSetVerificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verification",
		val,
	)
}

