package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for alert tags.
type AlertTags interface {
	IAlertTags
	Exclude() *[]*string
	SetExclude(val *[]*string)
	Include() *[]*string
	SetInclude(val *[]*string)
	Strength() *string
	SetStrength(val *string)
	Threshold() *string
	SetThreshold(val *string)
}

// The jsii proxy struct for AlertTags
type jsiiProxy_AlertTags struct {
	jsiiProxy_IAlertTags
}

func (j *jsiiProxy_AlertTags) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTags) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTags) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTags) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}


// Creates an instance of AlertTags.
func NewAlertTags(options IAlertTags) AlertTags {
	_init_.Initialize()

	if err := validateNewAlertTagsParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlertTags{}

	_jsii_.Create(
		"zap-cdk.AlertTags",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AlertTags.
func NewAlertTags_Override(a AlertTags, options IAlertTags) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AlertTags",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AlertTags)SetExclude(val *[]*string) {
	if err := j.validateSetExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_AlertTags)SetInclude(val *[]*string) {
	if err := j.validateSetIncludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_AlertTags)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_AlertTags)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

