package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the configuration for alert tags.
type AlertTag interface {
	IAlertTag
	Exclude() *[]*string
	SetExclude(val *[]*string)
	Include() *[]*string
	SetInclude(val *[]*string)
	Strength() *string
	SetStrength(val *string)
	Threshold() *string
	SetThreshold(val *string)
}

// The jsii proxy struct for AlertTag
type jsiiProxy_AlertTag struct {
	jsiiProxy_IAlertTag
}

func (j *jsiiProxy_AlertTag) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTag) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTag) Strength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlertTag) Threshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}


// Creates an instance of AlertTag.
func NewAlertTag(options IAlertTag) AlertTag {
	_init_.Initialize()

	j := jsiiProxy_AlertTag{}

	_jsii_.Create(
		"zap-cdk.AlertTag",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of AlertTag.
func NewAlertTag_Override(a AlertTag, options IAlertTag) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.AlertTag",
		[]interface{}{options},
		a,
	)
}

func (j *jsiiProxy_AlertTag)SetExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_AlertTag)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

func (j *jsiiProxy_AlertTag)SetStrength(val *string) {
	_jsii_.Set(
		j,
		"strength",
		val,
	)
}

func (j *jsiiProxy_AlertTag)SetThreshold(val *string) {
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

