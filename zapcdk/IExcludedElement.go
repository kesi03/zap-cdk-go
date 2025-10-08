package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing an excluded HTML element configuration.
type IExcludedElement interface {
	AttributeName() *string
	SetAttributeName(a *string)
	AttributeValue() *string
	SetAttributeValue(a *string)
	Description() *string
	SetDescription(d *string)
	Element() *string
	SetElement(e *string)
	Text() *string
	SetText(t *string)
	Xpath() *string
	SetXpath(x *string)
}

// The jsii proxy for IExcludedElement
type jsiiProxy_IExcludedElement struct {
	_ byte // padding
}

func (j *jsiiProxy_IExcludedElement) AttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExcludedElement)SetAttributeName(val *string) {
	_jsii_.Set(
		j,
		"attributeName",
		val,
	)
}

func (j *jsiiProxy_IExcludedElement) AttributeValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExcludedElement)SetAttributeValue(val *string) {
	_jsii_.Set(
		j,
		"attributeValue",
		val,
	)
}

func (j *jsiiProxy_IExcludedElement) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExcludedElement)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IExcludedElement) Element() *string {
	var returns *string
	_jsii_.Get(
		j,
		"element",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExcludedElement)SetElement(val *string) {
	if err := j.validateSetElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"element",
		val,
	)
}

func (j *jsiiProxy_IExcludedElement) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExcludedElement)SetText(val *string) {
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_IExcludedElement) Xpath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xpath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExcludedElement)SetXpath(val *string) {
	_jsii_.Set(
		j,
		"xpath",
		val,
	)
}

