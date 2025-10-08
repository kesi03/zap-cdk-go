package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IContextUser interface {
	Credentials() *[]IUserCredentials
	SetCredentials(c *[]IUserCredentials)
	Name() *string
	SetName(n *string)
}

// The jsii proxy for IContextUser
type jsiiProxy_IContextUser struct {
	_ byte // padding
}

func (j *jsiiProxy_IContextUser) Credentials() *[]IUserCredentials {
	var returns *[]IUserCredentials
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextUser)SetCredentials(val *[]IUserCredentials) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_IContextUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContextUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

