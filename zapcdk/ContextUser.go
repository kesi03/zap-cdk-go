package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents a user in the context.
type ContextUser interface {
	IContextUser
	Credentials() *[]IUserCredentials
	SetCredentials(val *[]IUserCredentials)
	Name() *string
	SetName(val *string)
}

// The jsii proxy struct for ContextUser
type jsiiProxy_ContextUser struct {
	jsiiProxy_IContextUser
}

func (j *jsiiProxy_ContextUser) Credentials() *[]IUserCredentials {
	var returns *[]IUserCredentials
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContextUser) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


// Creates an instance of ContextUser.
func NewContextUser(options IContextUser) ContextUser {
	_init_.Initialize()

	if err := validateNewContextUserParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContextUser{}

	_jsii_.Create(
		"zap-cdk.ContextUser",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ContextUser.
func NewContextUser_Override(c ContextUser, options IContextUser) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ContextUser",
		[]interface{}{options},
		c,
	)
}

func (j *jsiiProxy_ContextUser)SetCredentials(val *[]IUserCredentials) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_ContextUser)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

