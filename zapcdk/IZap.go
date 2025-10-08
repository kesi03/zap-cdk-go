package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing the ZAP (Zed Attack Proxy) configuration.
type IZap interface {
	Env() IEnvironment
	SetEnv(e IEnvironment)
	Jobs() *[]interface{}
	SetJobs(j *[]interface{})
}

// The jsii proxy for IZap
type jsiiProxy_IZap struct {
	_ byte // padding
}

func (j *jsiiProxy_IZap) Env() IEnvironment {
	var returns IEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IZap)SetEnv(val IEnvironment) {
	if err := j.validateSetEnvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"env",
		val,
	)
}

func (j *jsiiProxy_IZap) Jobs() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"jobs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IZap)SetJobs(val *[]interface{}) {
	if err := j.validateSetJobsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobs",
		val,
	)
}

