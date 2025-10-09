package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing an excluded HTML element configuration.
//
// Example:
//   const excludedElement = new ExcludedElement({
//     description: 'Exclude login button',
//     element: 'button',
//     text: 'Login'
//   });
//
type ExcludedElement interface {
	IExcludedElement
	AttributeName() *string
	SetAttributeName(val *string)
	AttributeValue() *string
	SetAttributeValue(val *string)
	Description() *string
	SetDescription(val *string)
	Element() *string
	SetElement(val *string)
	Text() *string
	SetText(val *string)
	Xpath() *string
	SetXpath(val *string)
}

// The jsii proxy struct for ExcludedElement
type jsiiProxy_ExcludedElement struct {
	jsiiProxy_IExcludedElement
}

func (j *jsiiProxy_ExcludedElement) AttributeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExcludedElement) AttributeValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExcludedElement) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExcludedElement) Element() *string {
	var returns *string
	_jsii_.Get(
		j,
		"element",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExcludedElement) Text() *string {
	var returns *string
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExcludedElement) Xpath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xpath",
		&returns,
	)
	return returns
}


// Creates an instance of ExcludedElement.
//
// Example:
//   const excludedElement = new ExcludedElement({
//     description: 'Exclude login button',
//     element: 'button',
//     text: 'Login'
//   });
//
func NewExcludedElement(options IExcludedElement) ExcludedElement {
	_init_.Initialize()

	if err := validateNewExcludedElementParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExcludedElement{}

	_jsii_.Create(
		"zap-cdk.ExcludedElement",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of ExcludedElement.
//
// Example:
//   const excludedElement = new ExcludedElement({
//     description: 'Exclude login button',
//     element: 'button',
//     text: 'Login'
//   });
//
func NewExcludedElement_Override(e ExcludedElement, options IExcludedElement) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.ExcludedElement",
		[]interface{}{options},
		e,
	)
}

func (j *jsiiProxy_ExcludedElement)SetAttributeName(val *string) {
	_jsii_.Set(
		j,
		"attributeName",
		val,
	)
}

func (j *jsiiProxy_ExcludedElement)SetAttributeValue(val *string) {
	_jsii_.Set(
		j,
		"attributeValue",
		val,
	)
}

func (j *jsiiProxy_ExcludedElement)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ExcludedElement)SetElement(val *string) {
	if err := j.validateSetElementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"element",
		val,
	)
}

func (j *jsiiProxy_ExcludedElement)SetText(val *string) {
	_jsii_.Set(
		j,
		"text",
		val,
	)
}

func (j *jsiiProxy_ExcludedElement)SetXpath(val *string) {
	_jsii_.Set(
		j,
		"xpath",
		val,
	)
}

