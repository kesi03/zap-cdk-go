package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the parameters for session management in the scanning process.
type SessionManagementParameters interface {
	ISessionManagementParameters
	Method() *string
	SetMethod(val *string)
	Parameters() ISessionManagementParametersParameters
	SetParameters(val ISessionManagementParametersParameters)
}

// The jsii proxy struct for SessionManagementParameters
type jsiiProxy_SessionManagementParameters struct {
	jsiiProxy_ISessionManagementParameters
}

func (j *jsiiProxy_SessionManagementParameters) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SessionManagementParameters) Parameters() ISessionManagementParametersParameters {
	var returns ISessionManagementParametersParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}


// Creates an instance of SessionManagementParameters.
func NewSessionManagementParameters(options ISessionManagementParameters) SessionManagementParameters {
	_init_.Initialize()

	if err := validateNewSessionManagementParametersParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_SessionManagementParameters{}

	_jsii_.Create(
		"zap-cdk.SessionManagementParameters",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of SessionManagementParameters.
func NewSessionManagementParameters_Override(s SessionManagementParameters, options ISessionManagementParameters) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.SessionManagementParameters",
		[]interface{}{options},
		s,
	)
}

func (j *jsiiProxy_SessionManagementParameters)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_SessionManagementParameters)SetParameters(val ISessionManagementParametersParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

