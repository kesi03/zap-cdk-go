package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ITechnology interface {
	Exclude() *[]*string
	SetExclude(e *[]*string)
	Include() *[]*string
	SetInclude(i *[]*string)
}

// The jsii proxy for ITechnology
type jsiiProxy_ITechnology struct {
	_ byte // padding
}

func (j *jsiiProxy_ITechnology) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITechnology)SetExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_ITechnology) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITechnology)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

