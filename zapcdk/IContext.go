package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IContext interface {
	Authentication() IAuthenticationParameters
	SetAuthentication(a IAuthenticationParameters)
	ExcludePaths() *[]*string
	SetExcludePaths(e *[]*string)
	IncludePaths() *[]*string
	SetIncludePaths(i *[]*string)
	Name() *string
	SetName(n *string)
	SessionManagement() ISessionManagementParameters
	SetSessionManagement(s ISessionManagementParameters)
	Structure() IContextStructure
	SetStructure(s IContextStructure)
	Technology() ITechnology
	SetTechnology(t ITechnology)
	Urls() *[]*string
	SetUrls(u *[]*string)
	Users() *[]IContextUser
	SetUsers(u *[]IContextUser)
}

// The jsii proxy for IContext
type jsiiProxy_IContext struct {
	_ byte // padding
}

func (j *jsiiProxy_IContext) Authentication() IAuthenticationParameters {
	var returns IAuthenticationParameters
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetAuthentication(val IAuthenticationParameters) {
	if err := j.validateSetAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_IContext) ExcludePaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetExcludePaths(val *[]*string) {
	_jsii_.Set(
		j,
		"excludePaths",
		val,
	)
}

func (j *jsiiProxy_IContext) IncludePaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetIncludePaths(val *[]*string) {
	_jsii_.Set(
		j,
		"includePaths",
		val,
	)
}

func (j *jsiiProxy_IContext) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IContext) SessionManagement() ISessionManagementParameters {
	var returns ISessionManagementParameters
	_jsii_.Get(
		j,
		"sessionManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetSessionManagement(val ISessionManagementParameters) {
	if err := j.validateSetSessionManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionManagement",
		val,
	)
}

func (j *jsiiProxy_IContext) Structure() IContextStructure {
	var returns IContextStructure
	_jsii_.Get(
		j,
		"structure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetStructure(val IContextStructure) {
	if err := j.validateSetStructureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"structure",
		val,
	)
}

func (j *jsiiProxy_IContext) Technology() ITechnology {
	var returns ITechnology
	_jsii_.Get(
		j,
		"technology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetTechnology(val ITechnology) {
	if err := j.validateSetTechnologyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"technology",
		val,
	)
}

func (j *jsiiProxy_IContext) Urls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetUrls(val *[]*string) {
	if err := j.validateSetUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urls",
		val,
	)
}

func (j *jsiiProxy_IContext) Users() *[]IContextUser {
	var returns *[]IContextUser
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContext)SetUsers(val *[]IContextUser) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

