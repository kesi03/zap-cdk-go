package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the credentials for a user.
type UserCredentials interface {
	IUserCredentials
	Password() *string
	SetPassword(val *string)
	Totp() ITotpConfig
	SetTotp(val ITotpConfig)
	Username() *string
	SetUsername(val *string)
}

// The jsii proxy struct for UserCredentials
type jsiiProxy_UserCredentials struct {
	jsiiProxy_IUserCredentials
}

func (j *jsiiProxy_UserCredentials) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserCredentials) Totp() ITotpConfig {
	var returns ITotpConfig
	_jsii_.Get(
		j,
		"totp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_UserCredentials) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}


// Creates an instance of UserCredentials.
func NewUserCredentials(options IUserCredentials) UserCredentials {
	_init_.Initialize()

	if err := validateNewUserCredentialsParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_UserCredentials{}

	_jsii_.Create(
		"zap-cdk.UserCredentials",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of UserCredentials.
func NewUserCredentials_Override(u UserCredentials, options IUserCredentials) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.UserCredentials",
		[]interface{}{options},
		u,
	)
}

func (j *jsiiProxy_UserCredentials)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_UserCredentials)SetTotp(val ITotpConfig) {
	_jsii_.Set(
		j,
		"totp",
		val,
	)
}

func (j *jsiiProxy_UserCredentials)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

