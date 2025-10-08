package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IUserCredentials interface {
	Password() *string
	SetPassword(p *string)
	Totp() ITotpConfig
	SetTotp(t ITotpConfig)
	Username() *string
	SetUsername(u *string)
}

// The jsii proxy for IUserCredentials
type jsiiProxy_IUserCredentials struct {
	_ byte // padding
}

func (j *jsiiProxy_IUserCredentials) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserCredentials)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_IUserCredentials) Totp() ITotpConfig {
	var returns ITotpConfig
	_jsii_.Get(
		j,
		"totp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserCredentials)SetTotp(val ITotpConfig) {
	_jsii_.Set(
		j,
		"totp",
		val,
	)
}

func (j *jsiiProxy_IUserCredentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IUserCredentials)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

