package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IPostman interface {
	CollectionFile() *string
	SetCollectionFile(c *string)
	CollectionUrl() *string
	SetCollectionUrl(c *string)
	Variables() *string
	SetVariables(v *string)
}

// The jsii proxy for IPostman
type jsiiProxy_IPostman struct {
	_ byte // padding
}

func (j *jsiiProxy_IPostman) CollectionFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostman)SetCollectionFile(val *string) {
	_jsii_.Set(
		j,
		"collectionFile",
		val,
	)
}

func (j *jsiiProxy_IPostman) CollectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostman)SetCollectionUrl(val *string) {
	_jsii_.Set(
		j,
		"collectionUrl",
		val,
	)
}

func (j *jsiiProxy_IPostman) Variables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPostman)SetVariables(val *string) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

