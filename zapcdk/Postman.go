package zapcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/kesi03/zap-cdk-go/zapcdk/jsii"
)

// Class representing the Postman import configuration.
//
// Example:
//   const postmanConfig = new Postman({
//     collectionFile: 'postman-collection.json',
//     variables: 'baseUrl=https://api.example.com,apiKey=12345'
//   });
//
type Postman interface {
	IPostman
	CollectionFile() *string
	SetCollectionFile(val *string)
	CollectionUrl() *string
	SetCollectionUrl(val *string)
	Variables() *string
	SetVariables(val *string)
}

// The jsii proxy struct for Postman
type jsiiProxy_Postman struct {
	jsiiProxy_IPostman
}

func (j *jsiiProxy_Postman) CollectionFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postman) CollectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Postman) Variables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


// Creates an instance of Postman.
func NewPostman(options IPostman) Postman {
	_init_.Initialize()

	if err := validateNewPostmanParameters(options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Postman{}

	_jsii_.Create(
		"zap-cdk.Postman",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Creates an instance of Postman.
func NewPostman_Override(p Postman, options IPostman) {
	_init_.Initialize()

	_jsii_.Create(
		"zap-cdk.Postman",
		[]interface{}{options},
		p,
	)
}

func (j *jsiiProxy_Postman)SetCollectionFile(val *string) {
	_jsii_.Set(
		j,
		"collectionFile",
		val,
	)
}

func (j *jsiiProxy_Postman)SetCollectionUrl(val *string) {
	_jsii_.Set(
		j,
		"collectionUrl",
		val,
	)
}

func (j *jsiiProxy_Postman)SetVariables(val *string) {
	_jsii_.Set(
		j,
		"variables",
		val,
	)
}

