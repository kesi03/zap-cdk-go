package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IEnvironmentProxy interface {
	Hostname() *string
	SetHostname(h *string)
	Password() *string
	SetPassword(p *string)
	Port() *float64
	SetPort(p *float64)
	Realm() *string
	SetRealm(r *string)
	Username() *string
	SetUsername(u *string)
}

// The jsii proxy for IEnvironmentProxy
type jsiiProxy_IEnvironmentProxy struct {
	_ byte // padding
}

func (j *jsiiProxy_IEnvironmentProxy) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProxy)SetHostname(val *string) {
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentProxy) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProxy)SetPassword(val *string) {
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentProxy) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProxy)SetPort(val *float64) {
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentProxy) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProxy)SetRealm(val *string) {
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_IEnvironmentProxy) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentProxy)SetUsername(val *string) {
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

