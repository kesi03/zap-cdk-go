package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IAlertTags interface {
	Exclude() *[]*string
	SetExclude(e *[]*string)
	Include() *[]*string
	SetInclude(i *[]*string)
	Strength() *string
	SetStrength(s *string)
	Threshold() *string
	SetThreshold(t *string)
}

// The jsii proxy for IAlertTags
type jsiiProxy_IAlertTags struct {
	_ byte // padding
}

func (j *jsiiProxy_IAlertTags) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTags)SetExclude(val *[]*string) {
	if err := j.validateSetExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_IAlertTags) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTags)SetInclude(val *[]*string) {
	if err := j.validateSetIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_IAlertTags) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTags)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_IAlertTags) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAlertTags)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

