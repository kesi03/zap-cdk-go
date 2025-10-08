package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IEnvironment interface {
	Contexts() *[]IContext
	SetContexts(c *[]IContext)
	Parameters() IEnvironmentParameters
	SetParameters(p IEnvironmentParameters)
	Proxy() IEnvironmentProxy
	SetProxy(p IEnvironmentProxy)
	Vars() *map[string]*string
	SetVars(v *map[string]*string)
}

// The jsii proxy for IEnvironment
type jsiiProxy_IEnvironment struct {
	_ byte // padding
}

func (j *jsiiProxy_IEnvironment) Contexts() *[]IContext {
	var returns *[]IContext
	_jsii_.Get(
		j,
		"contexts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment)SetContexts(val *[]IContext) {
	if err := j.validateSetContextsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contexts",
		val,
	)
}

func (j *jsiiProxy_IEnvironment) Parameters() IEnvironmentParameters {
	var returns IEnvironmentParameters
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment)SetParameters(val IEnvironmentParameters) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_IEnvironment) Proxy() IEnvironmentProxy {
	var returns IEnvironmentProxy
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment)SetProxy(val IEnvironmentProxy) {
	_jsii_.Set(
		j,
		"proxy",
		val,
	)
}

func (j *jsiiProxy_IEnvironment) Vars() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"vars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironment)SetVars(val *map[string]*string) {
	_jsii_.Set(
		j,
		"vars",
		val,
	)
}

