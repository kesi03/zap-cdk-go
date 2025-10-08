package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISessionManagementParameters interface {
	Method() *string
	SetMethod(m *string)
	Parameters() ISessionManagementParametersParameters
	SetParameters(p ISessionManagementParametersParameters)
}

// The jsii proxy for ISessionManagementParameters
type jsiiProxy_ISessionManagementParameters struct {
	_ byte // padding
}

func (j *jsiiProxy_ISessionManagementParameters) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionManagementParameters)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_ISessionManagementParameters) Parameters() ISessionManagementParametersParameters {
	var returns ISessionManagementParametersParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISessionManagementParameters)SetParameters(val ISessionManagementParametersParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

