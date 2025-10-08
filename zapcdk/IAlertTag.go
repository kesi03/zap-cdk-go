package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAlertTag interface {
	Exclude() *[]*string
	SetExclude(e *[]*string)
	Include() *[]*string
	SetInclude(i *[]*string)
	Strength() *string
	SetStrength(s *string)
	Threshold() *string
	SetThreshold(t *string)
}

// The jsii proxy for IAlertTag
type jsiiProxy_IAlertTag struct {
	_ byte // padding
}

func (j *jsiiProxy_IAlertTag) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTag)SetExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_IAlertTag) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTag)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_IAlertTag) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTag)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_IAlertTag) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTag)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

