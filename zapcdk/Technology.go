package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Represents the technology details for the scanning context.
type Technology interface {
	ITechnology
	Exclude() *[]*string
	SetExclude(val *[]*string)
	Include() *[]*string
	SetInclude(val *[]*string)
}

// The jsii proxy struct for Technology
type jsiiProxy_Technology struct {
	jsiiProxy_ITechnology
}

func (j *jsiiProxy_Technology) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Technology) Include() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}


// Creates an instance of Technology.
func NewTechnology(options ITechnology) Technology {
	_init_.Initialize()

	if err := validateNewTechnologyParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Technology{}

	_jsii_.Create(
		"zap-cdk.Technology",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Technology.
func NewTechnology_Override(t Technology, options ITechnology) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Technology",
		[]interface{}{options},
		t,
	)
}

func (j *jsiiProxy_Technology)SetExclude(val *[]*string) {
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_Technology)SetInclude(val *[]*string) {
	_jsii_.Set(
		j,
		"include",
		val,
	)
}

